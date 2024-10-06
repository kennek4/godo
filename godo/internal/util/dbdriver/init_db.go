package dbdriver

import (
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbDir *string) (err error) {

	// Get DB
	database, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	// Close connection when this function is done
	defer database.Close()

	const defaultTable string = "Tasks"
	
	err = CreateTableInDB(defaultTable, dbDir)
	if err != nil {
		return err
	}

	return nil
}
