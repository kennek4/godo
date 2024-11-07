package gddb

import (
	"fmt"
)

// Psuedo Enum Type
type DeleteType uint8

const (
	Title DeleteType = 1
	Id    DeleteType = 2
)

func DeleteGroup(groupName string) error {
	return nil
}

func Delete(deleteType DeleteType, group string, taskProperties []string) error {

	if GodoDb == nil {
		return ErrDbNotInitialized
	}

	if len(taskProperties) == 0 {
		return ErrInvalidDelete
	}

	var query string
	switch deleteType {
	case Title:
		query = fmt.Sprintf(`DELETE FROM "%s" WHERE title = ?`, group)
	case Id:
		query = fmt.Sprintf(`DELETE FROM "%s" WHERE id = ?`, group)
	}

	stmt, err := GodoDb.Prepare(query)
	if err != nil {
		return fmt.Errorf("%s, %s", ErrInvalidDelete, err)
	}

	for index := range taskProperties {
		_, err = stmt.Exec(taskProperties[index])
		if err != nil {
			return fmt.Errorf("%s, %s", ErrInvalidDelete, err)
		}
	}

	stmt.Close()
	return nil // Task(s) successfully deleted
}
