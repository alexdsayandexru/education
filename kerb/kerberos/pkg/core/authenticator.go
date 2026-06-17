package core

import (
	"encoding/json"
	"kerberos/pkg/sdk/http/http"
	"time"
)

func NewAuthenticator(secretKey32 []byte, crypto Crypto) (_authenticator []byte, err error) {
	authJson, _ := json.Marshal(Authenticator{
		Timestamp: time.Now(),
	})

	_authenticator, err = crypto.Encrypt(authJson, secretKey32)
	return
}

type Authenticator struct {
	Timestamp time.Time
	ClientId  []byte
}

func (t *Authenticator) Validate(lifetime time.Duration) error {
	if t.Timestamp.Add(lifetime).Unix() < time.Now().Unix() {
		return NewKerbError("authenticator is expired", http.ExpiredAuth)
	}
	return nil
}
