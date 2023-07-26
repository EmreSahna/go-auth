package handler

import (
	"encoding/base64"
	"encoding/json"
	"go-auth/http-basic-auth/model"
	"net/http"
	"os"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest model.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		return
	}

	if loginRequest.Username != os.Getenv("HTTP_BASIC_USERNAME") || loginRequest.Password != os.Getenv("HTTP_BASIC_PASSWORD") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	response := &model.LoginResponse{
		Message: "Successfully logged in",
	}

	encodedBasicAuth := base64.StdEncoding.EncodeToString([]byte(loginRequest.Username + ":" + loginRequest.Password))

	res, _ := json.Marshal(response)
	w.Header().Set("Authorization", "Basic "+encodedBasicAuth)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func SecretPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Secret Page"))
}
