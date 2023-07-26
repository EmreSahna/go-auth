package basicauth

import (
	"encoding/base64"
	"go-auth/http-basic-auth/model"
	"net/http"
	"os"
	"strings"
)

func DecodeBasicAuth(encodedAuth string) (model.User, error) {
	decodedAuth, err := base64.StdEncoding.DecodeString(encodedAuth)
	if err != nil {
		return model.User{}, err
	}

	auth := string(decodedAuth)
	credentialList := strings.Split(auth, ":")

	return model.User{
		Username: credentialList[0],
		Password: credentialList[1],
	}, nil
}

func CheckAuth(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authParam := r.Header.Get("Authorization")
		encodedAuth := authParam[len("Basic "):]
		if authParam != "" && encodedAuth != "" {
			decodedAuth, err := DecodeBasicAuth(encodedAuth)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if decodedAuth.Username == os.Getenv("HTTP_BASIC_USERNAME") && decodedAuth.Password == os.Getenv("HTTP_BASIC_PASSWORD") {
				handle(w, r)
				return
			}
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
