package dbdriver

import "fmt"

func InsertTaskInDB(title *string, description *string, table *string, dbDir *string) error {

	if title == nil || description == nil || table == nil {
		err := fmt.Errorf("in InsertTaskInDB, a supplied argument is a nil string pointer")
		return err
	}

	db, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	defer db.Close()

	query := fmt.Sprintf("INSERT INTO %s (title, description) VALUES (?, ?)", *table)
	statement, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("in InsertTaskInDB, something went wrong while preparing the query")
	}

	_, err = statement.Exec(*title, *description)
	if err != nil {
		return fmt.Errorf("in InsertTaskInDB, something went wrong while executing the query")
	}

	// Successfully inserted task into DB
	return nil
}
