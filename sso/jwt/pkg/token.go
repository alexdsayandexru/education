package pkg

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"time"
)

func GetTokenHS256(jti string, exp time.Time, secret string) (string, error) {
	claims := jwt.MapClaims{
		"jti": jti,
		"exp": exp.Unix(),
	}
	return NewTokenHS256(secret, claims)
}

func NewTokenHS256(secret string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateTokenHS256(ss string, secret string) (bool, error) {
	token, err := jwt.Parse(ss, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}
	return token.Valid, err
}

func ValidateTokenHS256WithClaims(ss string, secret string, claims jwt.Claims) (bool, error) {
	token, err := jwt.ParseWithClaims(ss, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return false, err
	}
	return token.Valid, err
}

func NewTokenRS256(privateKey []byte, claims jwt.Claims) (string, error) {
	pkey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(pkey)
}

func ValidateTokenRS256(ss string, publicKey []byte) (bool, error) {
	pkey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Fatalln(err)
	}

	token, err := jwt.Parse(ss, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}
		return pkey, nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, err
}

func ValidateTokenRS256WithClaims(ss string, publicKey []byte, claims jwt.Claims) (bool, error) {
	pkey, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Fatalln(err)
	}
	token, err := jwt.ParseWithClaims(ss, claims, func(token *jwt.Token) (interface{}, error) {
		return pkey, nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, err
}
