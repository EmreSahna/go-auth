package main

import (
	"github.com/joho/godotenv"
	"go-auth/session-based-auth/handler"
	"go-auth/session-based-auth/session"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/secret-page", session.ValidateSessionID(handler.SecretPage))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
