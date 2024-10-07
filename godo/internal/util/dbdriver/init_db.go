package dbdriver

import (
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(defaultTable string, dbDir *string) (err error) {

	// Get DB
	database, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	// Close connection when this function is done
	defer database.Close()

	err = CreateTableInDB(defaultTable, dbDir)
	if err != nil {
		return err
	}

	return nil
}
