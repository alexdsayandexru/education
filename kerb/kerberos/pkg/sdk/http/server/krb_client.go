package server

import (
	"kerberos/pkg/core"
	"time"
)

func NewKrbClient(serverHash []byte, crypto core.Crypto, authLifetime time.Duration) *KrbClient {
	return &KrbClient{
		serverHash:   serverHash,
		crypto:       crypto,
		authLifetime: authLifetime,
	}
}

type KrbClient struct {
	serverHash   []byte
	crypto       core.Crypto
	authLifetime time.Duration
}

func (t *KrbClient) Validate(_auth []byte, _tgs []byte) error {
	var tgs core.TGS

	if err := core.DecryptUnmarshal(_tgs, t.serverHash, t.crypto, &tgs); err != nil {
		return err
	} else if err = tgs.Validate(); err != nil {
		return err
	}

	var auth core.Authenticator

	if err := core.DecryptUnmarshal(_auth, tgs.SK2, t.crypto, &auth); err != nil {
		return err
	} else if err = auth.Validate(t.authLifetime); err != nil {
		return err
	}

	return nil
}
