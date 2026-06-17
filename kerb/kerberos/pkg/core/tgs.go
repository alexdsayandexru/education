package core

import (
	"kerberos/pkg/sdk/http/http"
	"time"
)

func NewTGS(sk2 []byte, timestamp time.Time, lifetime time.Time, clientId []byte) *TGS {
	return &TGS{
		SK2:       sk2,
		Timestamp: timestamp,
		Lifetime:  lifetime,
		ClientId:  clientId,
	}
}

type TGS struct {
	SK2       []byte
	Timestamp time.Time
	Lifetime  time.Time
	ClientId  []byte
}

func (t *TGS) Validate() error {
	if t.Lifetime.Unix() < time.Now().Unix() {
		return NewKerbError("TGS is expired", http.ExpiredTgs)
	}
	return nil
}
