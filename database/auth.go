package database

import (
	"time"

	"github.com/suynep/compilare/types"
)

func InsertUser(user types.RegisterUser) (int64, error) {
	q := `INSERT INTO users (username, email, password, last_login) VALUES (?, ?, ?, ?);`

	res, err := db.Exec(q, user.Username, user.Email, user.Password, time.Now())

	// if sqliteErr, ok := err.(*sqlite3.Error); ok && sqliteErr.Code == sqlite3.ErrConstraint {
	// 	return 0, err
	// }
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}

	return id, nil
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
	q := `SELECT id, username, email, password FROM users WHERE username=?;`

	user := new(types.RegisterUser)

	res := db.QueryRow(q, username)

	err := res.Scan(&user.Id, &user.Username, &user.Email, &user.Password)

	if err != nil {
		return *user, err
	}

	return *user, nil
}

func GetSessionByKey(key string) (types.Session, error) {
	q := `SELECT session_key, u_id FROM sessions WHERE session_key=?;`

	session := new(types.Session)

	res := db.QueryRow(q, key)

	err := res.Scan(&session.SessionKey, &session.UserId)

	if err != nil {
		return *session, err
	}

	return *session, nil

}

func GetSessionByUserId(u_id int64) (types.Session, error) {
	q := `SELECT session_key, u_id FROM sessions WHERE u_id=?;`

	session := new(types.Session)

	res := db.QueryRow(q, u_id)

	err := res.Scan(&session.SessionKey, &session.UserId)

	if err != nil {
		return *session, err
	}

	return *session, nil

}

func GetUserById(id int64) (types.RegisterUser, error) {
	q := `SELECT id, email, username, password FROM users WHERE id=?;`

	user := new(types.RegisterUser)

	res := db.QueryRow(q, id)

	err := res.Scan(&user.Id, &user.Email, &user.Username, &user.Password)

	if err != nil {
		return *user, err
	}

	return *user, nil

}
