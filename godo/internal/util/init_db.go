package util

import (
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbDir string) (err error) {

	// Get DB
	database, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	// Close connection when this function is done
	defer database.Close()

	// Create a table in the db
	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS godo (id INTEGER PRIMARY KEY, title TEXT, description TEXT)")
	if err != nil {
		return err
	}

	statement.Exec()

	return nil
}
