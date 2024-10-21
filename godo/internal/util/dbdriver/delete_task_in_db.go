package dbdriver

import (
	"fmt"
)

// Custom Psuedo Enum
type DeleteType int

const (
	title = iota + 1
	id
)

func getQueryString(queryType DeleteType, dbTable string, thingToDelete *interface{}) string {

	query := fmt.Sprintf("DELETE FROM %s", dbTable)

	switch queryType {
	case title:
		query = fmt.Sprintf("%s WHERE title='%s'", query, *thingToDelete)
	case id:
		query = fmt.Sprintf("%s WHERE id=%d", query, *thingToDelete)
	}

	return query
}

func DeleteTaskInDB(queryType DeleteType, dbTable string, thingToDelete *interface{}, dbDir *string) error {

	if thingToDelete == nil {
		err := fmt.Errorf("in DeleteTaskInDB, thingToDelete was supplied a nil interface pointer")
		return err
	}

	db, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	defer db.Close()

	query := getQueryString(queryType, dbTable, thingToDelete)
	fmt.Printf("query: %v\n", query)

	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}

	statement.Exec() // Exec query

	return nil // Task successfully deleted
}
