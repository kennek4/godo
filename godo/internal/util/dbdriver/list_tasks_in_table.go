package dbdriver

import (
	"fmt"
)

type Task struct {
	Id          int
	Title       string
	Description string
	IsComplete  bool
}

func ListTasksInTable(tableName *string, dbDir *string) (tasks []Task, err error) {

	if tableName == nil {
		err := fmt.Errorf("in ListTasksInTable, tableName was supplied with a nil string pointer")
		return nil, err
	}

	db, err := GetDB(dbDir)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM %s", *tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	// Add all tasks from the query 
	for rows.Next() {
		task := Task{}
		rows.Scan(&task.Id, &task.Title, &task.Description, &task.IsComplete)
		tasks = append(tasks, task)
	}

	return tasks, nil // Successfully printed table
}
