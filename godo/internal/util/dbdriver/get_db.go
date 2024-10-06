package dbdriver

import (
	"database/sql"
	"fmt"
	"path/filepath"
)

func GetDB(dbDir *string) (database *sql.DB, err error) {

	if dbDir == nil {
		err := fmt.Errorf("in GetDB, dbDir was supplied a nil string pointer")
		return nil, err
	}

	dbPath := filepath.Join(*dbDir, "godo.db")

	database, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return database, nil
}
