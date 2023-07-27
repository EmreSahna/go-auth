package session

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go-auth/session-based-auth/model"
	"net/http"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))
var sessions = map[string]model.Session{}

func GenerateSessionId(user model.User) string {
	combinedData := fmt.Sprintf("%s-%s-%s", user.Username, user.Password, secretKey)
	hashedData := sha256.Sum256([]byte(combinedData))
	sessionID := hex.EncodeToString(hashedData[:])
	return sessionID
}

func ValidateSessionID(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionID, err := r.Cookie("session_id")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		session, ok := sessions[sessionID.Value]
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		if session.ExpireAt.Before(time.Now()) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		handler.ServeHTTP(w, r)
	}
}

func SaveSession(session model.Session, sessionId string) {
	sessions[sessionId] = session
}
