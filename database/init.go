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
		id INTEGER PRIMARY KEY UNIQUE NOT NULL,
		deleted TEXT,
		type TEXT,
		by TEXT,
		time INTEGER,
		text TEXT,
		dead TEXT,
		parent INTEGER,
		poll TEXT,
		kids TEXT,
		url TEXT,
		score INTEGER,
		title TEXT,
		parts TEXT,
		descendants INTEGER
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

		// Enable foreign keys
		// if _, err = db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		// 	initErr = err
		// 	return
		// }

		// Create all tables
		tables := []string{createPostsTable}
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
