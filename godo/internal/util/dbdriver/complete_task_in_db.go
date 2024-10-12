package dbdriver

import (
	"fmt"

	"github.com/kennek4/godo/cmd"
)

func CompleteTask(taskId *int, currentGroup *string, dbDir *string) error {

	// Check to see if pointers are nil
	if taskId == nil || currentGroup == nil || dbDir == nil {
		return fmt.Errorf("in CompleteTask, an argument was provided a nil pointer")
	}

	db, err := GetDB(&cmd.GodoDir)
	if err != nil {
		return err
	}

	defer db.Close()

	// Check if ID is in current table (group)
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM %s WHERE id=%d)",
		*currentGroup, *taskId)

	result, err := db.Query(query)
	if err != nil {
		return err
	}

	var exists bool
	if result.Next() {
		var val int
		result.Scan(&val)

		exists = val != 0
	}

	result.Close()

	switch exists {
	case true:
		// Update Task
		query = fmt.Sprintf(`UPDATE %s SET isComplete=1 WHERE id=%d`, *currentGroup, *taskId)
		_, err = db.Exec(query)
		if err != nil {
			return err
		}
	case false:
		return fmt.Errorf("the given id does not exist in the current db, please try again")
	}

	return nil
}
