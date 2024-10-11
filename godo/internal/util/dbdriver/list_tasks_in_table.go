package dbdriver

import (
	"fmt"
)

const (
	incomplete = ' '
	complete   = '✔'
)

func ListTasksInTable(tableName *string, dbDir *string) error {

	if tableName == nil {
		err := fmt.Errorf("in ListTasksInTable, tableName was supplied with a nil string pointer")
		return err
	}

	db, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM %s", *tableName)
	rows, err := db.Query(query)
	if err != nil {
		return err
	}

	var id int
	var title string
	var description string
	var isComplete bool
	for rows.Next() {
		rows.Scan(&id, &title, &description, &isComplete)

		var checkBoxStatus rune
		switch {
		case isComplete:
			checkBoxStatus = complete
		case !isComplete:
			checkBoxStatus = incomplete
		}
		fmt.Printf(`
%d.) %s [%c]
%s		
`, id, title, checkBoxStatus, description)
	}

	fmt.Println()

	return nil // Successfully printed table
}
