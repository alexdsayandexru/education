package core

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

func NewSessionKey() []byte {
	return GetHash32(time.Now().String())
}

func GetHash32(key string) []byte {
	h := sha256.New()
	h.Write([]byte(key))
	return h.Sum(nil)
}

type Crypto interface {
	Encrypt(plaintext, secretKey32 []byte) ([]byte, error)
	Decrypt(ciphertext, secretKey32 []byte) ([]byte, error)
}

func DecryptUnmarshal(_jsonObj []byte, secretKey32 []byte, crypto Crypto, jsonObj interface{}) error {
	if tgs, err := crypto.Decrypt(_jsonObj, secretKey32); err != nil {
		return err
	} else {
		if err = json.Unmarshal(tgs, jsonObj); err != nil {
			return err
		}
	}
	return nil
}
