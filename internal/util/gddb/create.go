package gddb

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateGroup(groupName string) error {

	if GodoDb == nil {
		return ErrDbNotInitialized
	}

	stmt, err := GodoDb.Prepare(`CREATE TABLE IF NOT EXISTS ? (
		id INTEGER NOT NULL PRIMARY KEY,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		isComplete BOOLEAN NOT NULL DEFAULT FALSE);`)

	if err != nil {
		return ErrFailedToCreateGroup
	}

	defer stmt.Close()

	_, err = stmt.Exec(groupName) // Creates table
	if err != nil {
		return ErrFailedToCreateGroup
	}

	return nil // Group created successfully
}

func CreateTask(title string, desc string, table *string) error {
	if GodoDb == nil {
		return ErrDbNotInitialized
	}

	if title == "" {
		return ErrNoTitleForTask
	}

	if table == nil {
		return ErrDbInvalidGroup
	}

	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES (?, ?)", *table)
	stmt, err := GodoDb.Prepare(query)
	if err != nil {
		return ErrFailedToCreateTask
	}

	_, err = stmt.Exec(title, desc)
	if err != nil {
		return ErrFailedToCreateTask
	}

	// Successfully inserted task into DB
	return nil
}
