package dbdriver

import (
	"fmt"
	"strconv"
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

	count, err := db.Exec("SELECT COUNT(*) from " + (*tableName))
	if err != nil {
		return err
	}

	rowCount, err := count.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("Task count: " + strconv.Itoa(int(rowCount)))

	query := fmt.Sprintf("SELECT id, title, description FROM %s", *tableName)
	rows, err := db.Query(query)
	if err != nil {
		return err
	}

	var id int
	var title string
	var description string
	for rows.Next() {
		rows.Scan(&id, &title, &description)
		fmt.Printf("[%d] %s\n%s\n", id, title, description)
	}

	return nil // Successfully printed table
}
