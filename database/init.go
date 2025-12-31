package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

const (
	path             = "./data.db"
	createPostsTable = `CREATE TABLE IF NOT EXISTS posts(
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		pid INTEGER NOT NULL,
		deleted TEXT,
		type TEXT,
		by TEXT,
		time INTEGER,
		text TEXT,
		dead TEXT,
		parent INTEGER,
		poll TEXT,
		kids TEXT,
		url TEXT NOT NULL,
		score INTEGER,
		title TEXT,
		parts TEXT,
		descendants INTEGER,
		data_type TEXT NOT NULL,
    	UNIQUE (pid, url, data_type)
	);` // data_type can be either of "b", "t", "n" (best, top, new) stories

	createAeonPostsTable = `CREATE TABLE IF NOT EXISTS aeonposts(
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
                title TEXT,
                link  TEXT UNIQUE,     
                creator TEXT,   
                published TEXT,   
                description TEXT
	);`

	createPsychePostsTable = `CREATE TABLE IF NOT EXISTS psycheposts(
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
                title TEXT,
                link  TEXT UNIQUE,     
                creator TEXT,   
                published TEXT,   
                description TEXT
	);`

	createUserTable = `CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		username TEXT UNIQUE,
		password TEXT,
		email TEXT UNIQUE,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP,
		last_login TEXT
	);`

	createSessionTable = `CREATE TABLE IF NOT EXISTS sessions(
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		u_id INTEGER,
		session_key TEXT,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP,

		FOREIGN KEY (u_id) REFERENCES users(id),

		UNIQUE (u_id, session_key)
	);`

	createHackernewsFavoritesTable = `CREATE TABLE IF NOT EXISTS hackernewsfavorites(
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		u_id INTEGER,
		p_id INTEGER,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP,

		FOREIGN KEY (u_id) REFERENCES users(id),
		FOREIGN KEY (p_id) REFERENCES posts(pid),

		UNIQUE (u_id, p_id)
	);`

	createAeonFavoritesTable = `CREATE TABLE IF NOT EXISTS aeonfavorites(
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		u_id INTEGER,
		p_id INTEGER,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP,

		FOREIGN KEY (u_id) REFERENCES users(id),
		FOREIGN KEY (p_id) REFERENCES aeonposts(id),

		UNIQUE (u_id, p_id)
	);`

	createPsycheFavoritesTable = `CREATE TABLE IF NOT EXISTS psychefavorites(
		id INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
		u_id INTEGER,
		p_id INTEGER,
		created_at TEXT DEFAULT CURRENT_TIMESTAMP,

		FOREIGN KEY (u_id) REFERENCES users(id),
		FOREIGN KEY (p_id) REFERENCES psycheposts(id)

		UNIQUE (u_id, p_id)
	);`
)

var (
	db   *sql.DB
	once sync.Once
)

func InitDB() error {
	var initErr error
	once.Do(func() {
		// Ensure directory
		if err := os.MkdirAll("./", 0755); err != nil {
			initErr = err
			return
		}

		var err error
		db, err = sql.Open("sqlite3", path)
		if err != nil {
			initErr = err
			return
		}

		// Create all tables
		tables := []string{createPostsTable, createAeonPostsTable, createPsychePostsTable, createUserTable, createSessionTable, createAeonFavoritesTable, createHackernewsFavoritesTable, createPsycheFavoritesTable}
		for _, q := range tables {
			if _, err = db.Exec(q); err != nil {
				initErr = fmt.Errorf("failed to create table: %w", err)
				return
			}
		}
	})
	return initErr
}

// MustInitDB panics on failure — convenient for main()
func MustInitDB() {
	if err := InitDB(); err != nil {
		panic(err)
	}
}

func Close() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
