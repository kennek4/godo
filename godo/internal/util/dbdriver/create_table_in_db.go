package dbdriver

import "fmt"

func CreateTableInDB(newTableName string, dbDir *string) error {

	db, err := GetDB(dbDir)
	if err != nil {
		return fmt.Errorf("in CreateTableInDB, something went wrong with getting the db")
	}

	defer db.Close()

	table := fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	id INTEGER NOT NULL PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT,
	isComplete BOOLEAN NOT NULL DEFAULT FALSE);`, newTableName)

	_, err = db.Exec(table)
	if err != nil {
		return fmt.Errorf("in CreateTableInDB, something went wrong with executing the query")
	}

	return nil // Table was created
}
