package sqlite

import (
	"database/sql"
	
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLiteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS jobs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		salary INTEGER,
		country TEXT,
		skills TEXT
	)`)

	if err != nil {
		return nil, err
	}

	return db, nil
}
