package main

import (
	"encoding/json"
	"errors"
	"kerberos/pkg/core"
	"time"
)

const (
	AUTHLifetime time.Duration = time.Second
	SKLifetime   time.Duration = time.Second
	TGTLifetime  time.Duration = time.Second * 7
	TGSLifetime  time.Duration = time.Second * 2
)

func ToJson(t interface{}) []byte {
	j, _ := json.Marshal(t)
	return j
}

type KDR interface {
	GetPasswordHash32(clientId string) []byte
}

func NewKDC(kdr KDR, kdcHash32 []byte, crypto core.Crypto) *KDC {
	return &KDC{
		kdr:       kdr,
		kdcHash32: kdcHash32,
		crypto:    crypto,
	}
}

type KDC struct {
	kdr       KDR
	kdcHash32 []byte
	crypto    core.Crypto
}

func (t *KDC) GetTGT(clientId []byte, _authenticator []byte) (_sk1 []byte, _tgt []byte, err error) {
	clientHash32 := t.kdr.GetPasswordHash32(string(clientId))
	if clientHash32 == nil {
		err = errors.New("not found client hash password")
		return
	}

	if err = t.authenticate(_authenticator, clientHash32); err != nil {
		return
	}

	sk1 := core.NewSessionKey()

	_sk1 = t._SK(sk1, clientHash32)
	_tgt = t._TGT(sk1, clientId)

	return
}

func (t *KDC) GetTGS(serverId []byte, _authenticator []byte, _tgt []byte) (_sk2 []byte, _tgs []byte, err error) {
	serverHash32 := t.kdr.GetPasswordHash32(string(serverId))
	if serverHash32 == nil {
		err = errors.New("not found server hash password")
		return
	}

	tgt, err := t.TGT(_tgt)
	if err != nil {
		return
	} else if err = tgt.Validate(); err != nil {
		return
	}

	if err = t.authenticate(_authenticator, tgt.SK1); err != nil {
		return
	}

	sk2 := core.NewSessionKey()

	_sk2 = t._SK(sk2, tgt.SK1)
	_tgs = t._TGS(sk2, serverId, serverHash32)

	return
}

func (t *KDC) authenticate(_auth, hash32 []byte) error {
	var auth core.Authenticator
	if err := core.DecryptUnmarshal(_auth, hash32, t.crypto, &auth); err != nil {
		return err
	}
	return auth.Validate(AUTHLifetime)
}

func (t *KDC) _SK(sessionKey []byte, hash32 []byte) (_sk []byte) {
	sk := core.NewSK(sessionKey, time.Now(), time.Now().Add(SKLifetime))
	_sk, _ = t.crypto.Encrypt(ToJson(sk), hash32)
	return
}

func (t *KDC) _TGT(sessionKey []byte, clientId []byte) (_tgt []byte) {
	tgt := core.NewTGT(sessionKey, time.Now(), time.Now().Add(TGTLifetime), clientId)
	_tgt, _ = t.crypto.Encrypt(ToJson(tgt), t.kdcHash32)
	return
}

func (t *KDC) _TGS(sessionKey, serverId, serverHash32 []byte) (_tgs []byte) {
	tgs := core.NewTGS(sessionKey, time.Now(), time.Now().Add(TGSLifetime), serverId)
	_tgs, _ = t.crypto.Encrypt(ToJson(tgs), serverHash32)
	return
}

func (t *KDC) TGT(_tgt []byte) (*core.TGT, error) {
	var tgt core.TGT
	if err := core.DecryptUnmarshal(_tgt, t.kdcHash32, t.crypto, &tgt); err != nil {
		return nil, err
	}
	return &tgt, nil
}
