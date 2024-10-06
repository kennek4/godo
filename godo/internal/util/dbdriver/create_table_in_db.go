package dbdriver

import "fmt"

func CreateTableInDB(newTableName *string, dbDir *string) error {

	if newTableName == nil {
		return fmt.Errorf("the given table name argument is nil")
	}

	db, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	table := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id INTEGER NOT NULL PRIMARY KEY, 
		title TEXT NOT NULL,
		description TEXT,
		isComplete BOOLEAN NOT NULL DEFAULT FALSE);`)

	return nil // Table was created
}
