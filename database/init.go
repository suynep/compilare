package database

import (
	_ "database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	PostTable = `CREATE TABLE IF NOT EXISTS posts(
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

// "type,omitempty" db:"ty	)
// "by,omitempty" db:"by,o)
// "time,omitempty" db:"ti
// "text,omitempty" db:"te
// "dead,omitempty" db:"de
// "parent,omitempty" db:"
// "poll,omitempty" db:"po
// "kids,omitempty" db:"ki
// "url,omitempty" db:"url
// "score,omitempty" db:"s
// "title,omitempty" db:"t
// "parts,omitempty" db:"p
// "descendants,omitempty"
