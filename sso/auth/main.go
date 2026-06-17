package main

import (
	"auth/entities"
	"auth/handlers"
	"auth/params"
	"auth/services"
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"strings"
)

const (
	ClientId      = "client_id"
	ClientSecret  = "client_secret"
	RefreshToken  = "refresh_token"
	Authorization = "Authorization"
	GrantType     = "grant_type"
)

func main() {
	http.HandleFunc("/oauth/token", token)
	http.HandleFunc("/oauth/validate", validate)

	log.Fatal(http.ListenAndServe(":8888", nil))
}

func auth(w http.ResponseWriter, r *http.Request) {
	/*grantType := r.URL.Query().Get(GrantType)
	clientId := r.URL.Query().Get(ClientId)
	clientSecret := r.URL.Query().Get(ClientSecret)
	refreshToken := r.URL.Query().Get(RefreshToken)*/
}

func validate(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header.Get(Authorization), " ")[1]

	var claims jwt.MapClaims
	_, _ = jwt.ParseWithClaims(token, &claims, nil)

	secret := services.Repo.Get(claims["jti"].(string))

	if ok, err := services.NewTokenService().Validate(token, secret); !ok {
		Response(w, entities.NewErrorResponse(http.StatusBadRequest, entities.InvalidRequest, err.Error()))
	}
}

func token(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		Response(w, entities.NewErrorResponse(http.StatusBadRequest, entities.InvalidRequest, err.Error()))
		return
	}

	grantType := r.Form.Get(GrantType)
	clientId := r.Form.Get(ClientId)
	clientSecret := r.Form.Get(ClientSecret)
	refreshToken := r.Form.Get(RefreshToken)

	if err := params.Valid().
		Required(GrantType, grantType).
		Required(ClientId, clientId).
		Required(ClientSecret, clientSecret).
		RequiredIf(RefreshToken, refreshToken, func() bool { return grantType == RefreshToken }).Result(); err != nil {
		Response(w, entities.NewErrorResponse(http.StatusBadRequest, entities.InvalidRequest, err.Error()))
	} else {

		handler := handlers.NewTokenEventHandler(grantType, clientId, clientSecret, refreshToken, services.NewClientService(),
			services.NewTokenService())

		if at, rt, err := handler.Handle(); err == nil {
			Response(w, entities.NewTokenResponse(at, rt))
		} else {
			Response(w, entities.NewErrorResponse(http.StatusBadRequest, entities.InvalidRequest, err.Error()))
		}
	}
}

func Response(w http.ResponseWriter, r interface{}) {
	t, _ := json.Marshal(r)
	if _, err := w.Write(t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
