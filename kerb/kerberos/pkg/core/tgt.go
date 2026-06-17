package core

import (
	"kerberos/pkg/sdk/http/http"
	"time"
)

func NewTGT(sk1 []byte, timestamp time.Time, lifetime time.Time, clientId []byte) *TGT {
	return &TGT{
		SK1:       sk1,
		Timestamp: timestamp,
		Lifetime:  lifetime,
		ClientId:  clientId,
	}
}

type TGT struct {
	SK1       []byte
	Timestamp time.Time
	Lifetime  time.Time
	ClientId  []byte
}

func (t *TGT) Validate() error {
	if t.Lifetime.Unix() < time.Now().Unix() {
		return NewKerbError("TGT is expired", http.ExpiredTgt)
	}
	return nil
}
