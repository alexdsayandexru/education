package services

import (
	"github.com/google/uuid"
	"jwt/pkg"
	"time"
)

type TokenService interface {
	Create(secret string, aTokenLifetime, rTokenLifetime time.Duration) (at string, rt string, err error)
	Refresh(refresh string, secret string) (at string, rt string, err error)
	Validate(token string, secret string) (bool, error)
}

func NewTokenService() TokenService {
	return &TokenServiceImpl{}
}

type TokenServiceImpl struct {
}

func (t *TokenServiceImpl) Create(secret string, aTokenLifetime, rTokenLifetime time.Duration) (string, string, error) {
	tokenId := uuid.New().String()
	Repo.Set(tokenId, secret)

	access, _ := pkg.GetTokenHS256(
		tokenId,
		time.Now().Add(time.Minute*aTokenLifetime),
		secret)

	refresh, _ := pkg.GetTokenHS256(
		tokenId,
		time.Now().Add(time.Minute*rTokenLifetime),
		secret)

	return access, refresh, nil
}

func (t *TokenServiceImpl) Refresh(refresh string, secret string) (string, string, error) {
	if ok, err := pkg.ValidateTokenHS256(refresh, secret); !ok {
		return "", "", err
	}

	tokenId := uuid.New().String()
	Repo.Set(tokenId, secret)

	access, _ := pkg.GetTokenHS256(
		tokenId,
		time.Now().Add(time.Minute*time.Duration(1)),
		secret)

	return access, refresh, nil
}

func (t *TokenServiceImpl) Validate(token string, secret string) (bool, error) {
	return pkg.ValidateTokenHS256(token, secret)
}
