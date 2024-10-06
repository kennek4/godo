package util

import (
	"fmt"

	"github.com/pelletier/go-toml/query"
)

// Custom Psuedo Enum
type DeleteType int

const (
	id = iota + 1
	title
	description
)

func getQueryString(queryType DeleteType, dbTable *string, thingToDelete *interface{}) string {

	query := fmt.Sprintf("DELETE FROM %s", *dbTable)

	switch queryType {
	case id:
		query = fmt.Sprintf("%s WHERE id=%d", query, *thingToDelete)
	case title:
		query = fmt.Sprintf("%s WHERE title='%s'", query, *thingToDelete)
	case description:
		query = fmt.Sprintf("%s WHERE description='%s'", query, *thingToDelete)

	}

	return query
}

func DeleteTaskInDB(queryType *int, dbTable *string, dbDir *string) error {

	db, err := GetDB(dbDir)
	if err != nil {
		return err
	}

	defer db.Close()

	query := getQueryString()

	return nil // Task successfully deleted
}
