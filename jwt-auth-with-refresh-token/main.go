package main

import (
	"github.com/joho/godotenv"
	"go-auth/jwt-auth-with-refresh-token/handler"
	"go-auth/jwt-auth-with-refresh-token/jwt"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/secret-page", jwt.ValidateJWT(handler.SecretPage))
	http.HandleFunc("/refresh", handler.Refresh)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
