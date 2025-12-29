package types

import (
	"time"
)

type RegisterUser struct {
	Username string `db:"username,omitempty"`
	Password string `db:"password,omitempty"`
	Email    string `db:"email,omitempty"`
}

type User struct {
	Id        int
	Username  string    `db:"username,omitempty"`
	Password  string    `db:"password,omitempty"`
	CreatedAt time.Time `db:"created_at,omitempty"`
	LastLogin time.Time `db:"last_login,omitempty"`
}

type Session struct {
	SessionKey string `db:"session_key" json:"session_key"`
	UserId     int    `db:"u_id" json:"user_id"`
}
