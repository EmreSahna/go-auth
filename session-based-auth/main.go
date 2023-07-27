package main

import (
	"go-auth/session-based-auth/handler"
	"go-auth/session-based-auth/session"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/secret-page", session.ValidateSessionID(handler.SecretPage))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
