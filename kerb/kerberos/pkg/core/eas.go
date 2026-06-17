package core

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

func NewEAS() Crypto {
	return &EAS{}
}

type EAS struct {
}

func (*EAS) Encrypt(plaintext, secretKey []byte) ([]byte, error) {
	_aes, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(_aes)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	return ciphertext, nil
}

func (*EAS) Decrypt(ciphertext, secretKey []byte) ([]byte, error) {
	_aes, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(_aes)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
