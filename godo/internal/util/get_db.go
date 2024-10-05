package util

import (
	"database/sql"
	"path/filepath"
)

func GetDB(dbDir string) (database *sql.DB, err error) {

	dbPath := filepath.Join(dbDir, "godo.db")

	database, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return database, nil
}
