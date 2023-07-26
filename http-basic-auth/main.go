package main

import (
	"github.com/joho/godotenv"
	"go-auth/http-basic-auth/basic-auth"
	"go-auth/http-basic-auth/handler"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/secret-page", basicauth.CheckAuth(handler.SecretPage))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
