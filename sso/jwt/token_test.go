package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"jwt/pkg"
	"os"
	"testing"
	"time"
)

/*
Issuer    string       `json:"iss,omitempty"`
Subject   string       `json:"sub,omitempty"`
Audience  ClaimStrings `json:"aud,omitempty"`
ExpiresAt *NumericDate `json:"exp,omitempty"`
NotBefore *NumericDate `json:"nbf,omitempty"`
IssuedAt  *NumericDate `json:"iat,omitempty"`
ID        string       `json:"jti,omitempty"`
*/

func NewMapClaims() jwt.MapClaims {
	/*return jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2024, 9, 16, 12, 0, 0, 0, time.UTC).Unix(),
		"exp": time.Date(2024, 9, 18, 12, 0, 0, 0, time.UTC).Unix(),
	}*/
	return jwt.MapClaims{
		"jti": "qwerty",
		"exp": time.Now().Add(time.Minute).Unix(),
	}
}

func NewMyCustomClaims() MyCustomClaims {
	expDateTime, _ := time.Parse("2006-01-02", "2024-09-19")
	isuDateTime, _ := time.Parse("2006-01-02", "2024-09-12")
	nbfDateTime, _ := time.Parse("2006-01-02", "2024-09-10")

	return MyCustomClaims{
		[]string{"user", "admin", "reader", "writer"},
		jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(expDateTime),
			IssuedAt:  jwt.NewNumericDate(isuDateTime),
			NotBefore: jwt.NewNumericDate(nbfDateTime),
			Issuer:    "test",
			Subject:   "somebody",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
	}
}

func TestNewTokenHS256(t *testing.T) {
	token, _ := pkg.NewTokenHS256("helloworld", NewMapClaims())
	fmt.Println(token)

	if ok, err := pkg.ValidateTokenHS256(token, "helloworld"); !ok {
		t.Error(err)
	}
}

type MyCustomClaims struct {
	Roles []string `json:"roles"`
	jwt.RegisteredClaims
}

func Test2NewTokenHS256(t *testing.T) {
	token, _ := pkg.NewTokenHS256("helloworld", NewMyCustomClaims())
	fmt.Println(token)

	claims := MyCustomClaims{}
	if ok, err := pkg.ValidateTokenHS256WithClaims(token, "helloworld", &claims); !ok {
		t.Error(err)
	} else {
		fmt.Println(claims.Roles)
	}
}

func TestTokenRSA256(t *testing.T) {
	pubKey, err := os.ReadFile("cert/skey.pub")
	if err != nil {
		t.Error(err)
	}
	token := "eyJhbGciOiJSUzI1NiIsImtpZCI6InB1YmxpYzo1MDcyYzE4Yy03ZTJlLTRiNGEtYWRiYS0wYzdhNTM0OWZmNTciLCJ0eXAiOiJKV1QifQ.eyJhdF9oYXNoIjoiY1ZndVBBdHNMUVNIcUVIU0NseUliUSIsImF1ZCI6WyJwZXJzb25hbC1hY2NvdW50Il0sImF1dGhfdGltZSI6MTcyOTA2NjM2OSwiZXhwIjoxNzI5MDY5OTcxLCJpYXQiOjE3MjkwNjYzNzEsImlzcyI6Imh0dHBzOi8vYXV0aC5naWQucnUvIiwianRpIjoiZTU5ZTdiYzctN2FkYy00YTVlLWJiYzgtOWM2NDY4MDg5MmQwIiwicmF0IjoxNzI5MDY2MzU1LCJzaWQiOiI3YjE1MWNmOS1iNWJlLTRiZDItYTFjNC0xNGI0NzAwMTA0OTkiLCJzdWIiOiJRU21Od3dTd3hnQ3RMbzh5aFhwa2JTIn0.cOCiP0QYMG2Z2cBm4ZpPS3Ti6OMGFlBcnjQ96_I_leHId1OAxm_gTxEljJnVp3fb-mOu5DwFjzW2ci0aHyFJXwh7vWJJO_QgE4DV6VinwlU8-eGTMnvjpwPDr5z8I3aj9qIJLDmcbo5uerCDOHrBYHR5Yv81k1C61WxhcCepN5Fm-bcD4BemoA_3l52B3EPDTynxXK84xOwAY4tdsrsK4hquDbXiOKL5B3Gu_2uRMKB-r7HIZ6Xsv598TakDYxUU1jwydVQInGh-GpASaXxJP9S-93X6xVVFd_lRNoqVqbc7Z-YHe869xSfjtzEa1UI79npS6lO4XImVWpur02YZmtEkiXogNKbN6CFSY90Y-ukVkRpAK9Oc-INBiP4XUY5miX7BXAk8Vl6fhZgEB5EghgQhk5R9JsKaFdKrwV17i0gV8d6HDNtYis-WHXrGVkuFyl6CNDT43uyyUDchwN7q2ueZm0FUHur6K8rgbFgHrmTjU5YtktEe74cKp_FJZp7XlTrPf6oOnAdfB_ocLrMWFiA-AajFemxyqujW0a1WawlQcNCLWRTaOyfznSpTUAR6n9Pkjk4UR7oHFmWYoKgzXnmPIlpDQMb97LeRPGM3xXo_BHLlOEGSHd7m7iSOdUJEwmd4hhkTfmIdURngvuFt_C2M5FS9r2_A2pGtRyk19Wo"
	if ok, err := pkg.ValidateTokenRS256(token, pubKey); !ok {
		t.Error(err)
	}
}

func TestNewTokenRSA256(t *testing.T) {
	prvKey, err := os.ReadFile("cert/id_rsa")
	if err != nil {
		t.Error(err)
	}

	token, err := pkg.NewTokenRS256(prvKey, NewMapClaims())
	if err != nil {
		t.Error(err)
	}

	pubKey, err := os.ReadFile("cert/id_rsa.pub")
	if err != nil {
		t.Error(err)
	}

	if ok, err := pkg.ValidateTokenRS256(token, pubKey); !ok {
		t.Error(err)
	}
}

func Test2NewTokenRSA256(t *testing.T) {
	prvKey, err := os.ReadFile("cert/id_rsa")
	if err != nil {
		t.Error(err)
	}

	token, err := pkg.NewTokenRS256(prvKey, NewMyCustomClaims())
	if err != nil {
		t.Error(err)
	}

	pubKey, err := os.ReadFile("cert/id_rsa.pub")
	if err != nil {
		t.Error(err)
	}

	claims := MyCustomClaims{}
	if ok, err := pkg.ValidateTokenRS256WithClaims(token, pubKey, &claims); !ok {
		t.Error(err)
	} else {
		fmt.Println(claims.Roles)
	}
}
