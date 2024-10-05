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

	const table string = `
	"CREATE TABLE IF NOT EXISTS godo (
		id INTEGER NOT NULL PRIMARY KEY, 
		title TEXT NOT NULL,
		description TEXT,
		isComplete BOOLEAN NOT NULL DEFAULT FALSE)"
	`

	// Create a table in the db
	statement, err := database.Prepare(table)
	if err != nil {
		return err
	}

	statement.Exec()

	return nil
}
