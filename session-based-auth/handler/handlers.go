package handler

import (
	"encoding/json"
	"go-auth/session-based-auth/model"
	"go-auth/session-based-auth/session"
	"net/http"
	"time"
)

var users = map[string]string{
	"user": "pass",
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	pass, ok := users[user.Username]
	if !ok || pass != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	expiresAt := time.Now().Add(time.Minute * 2)
	sessionID := session.GenerateSessionId(user)
	session.SaveSession(model.Session{
		Username: user.Username,
		ExpireAt: expiresAt,
	}, sessionID)

	http.SetCookie(w, &http.Cookie{
		Name:    "session_id",
		Value:   sessionID,
		Expires: expiresAt,
	})
	w.Write([]byte("Successfully logged in"))
}

func SecretPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Secret Page"))
}
