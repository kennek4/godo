package gddb

import "fmt"

type Group struct {
	Name      string
	TaskCount uint8
}

type Task struct {
	Id          int
	Title       string
	Description string
	IsComplete  bool
}

func GetGroup() (groups []Group, err error) {
	if GodoDb == nil {
		return nil, ErrDbNotInitialized
	}

	query := `SELECT tbl_name FROM sqlite_master WHERE type="table"`
	rows, err := GodoDb.Query(query)
	if err != nil {
		return nil, err
	}

	statement, err := GodoDb.Prepare("SELECT COUNT(*) FROM ?")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	for rows.Next() {
		group := Group{}
		rows.Scan(&group.Name)

		err = statement.QueryRow(group.Name).Scan(&group.TaskCount)
		if err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	return groups, nil
}

func GetTasks(group string) (tasks []Task, err error) {

	if GodoDb == nil {
		return nil, ErrDbNotInitialized
	}

	if group == "" {
		return nil, ErrDbInvalidGroup
	}

	cleansed := fmt.Sprintf(`SELECT * FROM "%s"`, group)
	stmt, err := GodoDb.Prepare(cleansed)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
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
