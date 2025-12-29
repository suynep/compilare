package database

import (
	"github.com/suynep/compilare/types"
	"time"
)

func InsertUser(user types.RegisterUser) int64 {
	q := `INSERT OR IGNORE INTO users (username, email, password, last_login) VALUES (?, ?, ?, ?);`

	res, err := db.Exec(q, user.Username, user.Email, user.Password, time.Now())

	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}

	return id
}

func InsertSession(session types.Session) int64 {
	q := `INSERT OR IGNORE INTO sessions (u_id, session_key) VALUES (?, ?);`

	res, err := db.Exec(q, session.UserId, session.SessionKey)

	if err != nil {
		panic(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}

	return id
}

func GetUserByUsername(username string) (types.RegisterUser, error) {
	q := `SELECT username, email, password FROM users WHERE username=?;`

	user := new(types.RegisterUser)

	res, err := db.Query(q, username)

	if err != nil {
		return *user, err
	}

	err = res.Scan(user)

	if err != nil {
		return *user, err
	}

	return *user, nil
}

func GetSessionByKey(key string) (types.Session, error) {
	q := `SELECT session_key, u_id FROM sessions WHERE session_key=?;`

	session := new(types.Session)

	res, err := db.Query(q, key)

	if err != nil {
		return *session, err
	}

	err = res.Scan(session)

	if err != nil {
		return *session, err
	}

	return *session, nil

}
