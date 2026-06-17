package handlers

import (
	"auth/entities"
	"auth/services"
)

const (
	ClientCredentials = "client_credentials"
	RefreshToken      = "refresh_token"
)

func NewTokenEventHandler(grantType string, clientId string, clientSecret string, refreshToken string,
	cs services.ClientService, ts services.TokenService) *TokenEventHandler {
	return &TokenEventHandler{
		grantType:    grantType,
		clientId:     clientId,
		clientSecret: clientSecret,
		refreshToken: refreshToken,
		cs:           cs,
		ts:           ts,
	}
}

type TokenEventHandler struct {
	grantType    string
	clientId     string
	clientSecret string
	refreshToken string
	cs           services.ClientService
	ts           services.TokenService
	error        *entities.ErrorResponse
}

func (t *TokenEventHandler) Handle() (string, string, error) {
	if t.grantType == ClientCredentials {
		return t.ts.Create(t.clientSecret, 1, 5)
	} else if t.grantType == RefreshToken {
		return t.ts.Refresh(t.refreshToken, t.clientSecret)
	} else {
		panic("Unexpected GrantType")
	}
}
