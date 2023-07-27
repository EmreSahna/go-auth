package model

import "time"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Session struct {
	Username string
	ExpireAt time.Time
}
