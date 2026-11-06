package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func New() *sql.DB {
	db, err := sql.Open("sqlite", "./db")
	if err != nil {
		return nil
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil
	}

	return db
}
