package entities

const Bearer = "Bearer"

func NewTokenResponse(accessToken, refreshToken string) *TokenResponse {
	return &TokenResponse{
		TokenType:    Bearer,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

type TokenResponse struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
