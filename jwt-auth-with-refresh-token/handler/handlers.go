package handler

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v5"
	jwtHelper "go-auth/jwt-auth-with-refresh-token/jwt"
	"go-auth/jwt-auth-with-refresh-token/model"
	"net/http"
	"time"
)

var users = map[string]string{
	"user": "123",
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		return
	}

	password, ok := users[credentials.Username]
	if !ok || password != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	accessClaims := &model.Claims{
		Username: credentials.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 10)),
		},
	}

	refreshClaims := &model.Claims{
		Username: credentials.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
		},
	}

	accessToken, err := jwtHelper.CreateJWT(accessClaims)
	refreshToken, err := jwtHelper.CreateJWT(refreshClaims)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	response := &model.LoginResponse{
		Message:      "Successfully logged in",
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	res, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func SecretPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Secret Page"))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	var refreshRequest model.RefreshRequest
	err := json.NewDecoder(r.Body).Decode(&refreshRequest)
	if err != nil {
		return
	}

	sentToken := refreshRequest.RefreshToken
	newAccessToken, _ := jwtHelper.CreateAccessTokenWithRefreshToken(sentToken)
	if newAccessToken == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	response := &model.RefreshResponse{
		Message:     "Successfully refreshed",
		AccessToken: newAccessToken,
	}
	res, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
