package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func NewSQLiteStorage(filename string) (*Storage, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}
