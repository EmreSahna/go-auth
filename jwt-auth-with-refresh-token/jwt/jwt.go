package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"go-auth/jwt-auth-with-refresh-token/model"
	"net/http"
	"os"
	"time"
)

func CreateJWT(claims *model.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func ValidateJWT(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		sentToken := auth[len("Bearer "):]
		if sentToken != "" {
			token, _ := jwt.Parse(sentToken, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET")), nil
			})

			if token.Valid {
				handler(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		}
	}
}

func CreateAccessTokenWithRefreshToken(refreshToken string) (string, error) {
	token, _ := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if token.Valid {
		claims := &model.Claims{
			Username: token.Claims.(jwt.MapClaims)["username"].(string),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 10)),
			},
		}

		return CreateJWT(claims)
	}
	return "", nil
}
