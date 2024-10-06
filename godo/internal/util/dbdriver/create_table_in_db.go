package dbdriver

import "fmt"

func CreateTableInDB(newTableName string, dbDir *string) error {

	db, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	defer db.Close()

	table := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id INTEGER NOT NULL PRIMARY KEY, 
		title TEXT NOT NULL,
		description TEXT,
		isComplete BOOLEAN NOT NULL DEFAULT FALSE);`, newTableName)

	statement, err := db.Prepare(table)
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil // Table was created
}
