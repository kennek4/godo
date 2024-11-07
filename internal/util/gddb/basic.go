package gddb

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"path/filepath"
)

var (
	ErrDbNotInitialized    = errors.New("a database for Godo has yet to be initialized")
	ErrDbInvalidGroup      = errors.New("an invalid group was supplied as an argument")
	ErrFailedToCreateGroup = errors.New("something went wrong while trying to create a Godo group")
	ErrFailedToCreateTask  = errors.New("something went wrong while trying to create a Godo task")
	ErrInvalidDelete       = errors.New("something went wrong while trying to delete a Godo task or group")
	ErrNoTitleForTask      = errors.New("no title was given when trying to create a task, a title is needed")
)

var GodoDb *sql.DB

func Load(godoPath string) error {
	dbPath := filepath.Join(godoPath, "godo.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("the godo db could not be opened, %s", err)
	}

	GodoDb = db
	return nil // db loaded successfully
}

func Close(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal("Something went catastrophically wrong when closing the GodoDB connection")
	}
}
