package dbdriver

import (
	"fmt"
	"os"
	"strings"

	"github.com/kennek4/godo/internal/util/console"
	"golang.org/x/term"
)

const (
	incomplete = ' '
	complete   = '✔'
)

func ListTasksInTable(tableName *string, dbDir *string) error {

	console.ClearConsole()

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

	table := strings.ToUpper(fmt.Sprintf("%s TASKS", *tableName))
	fmt.Println(table + "\n")

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}

	isFirstColumn := true // The "column" that the task will be printed in

	var line []rune
	for range width {
		line = append(line, '-')
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

		switch isFirstColumn {
		case true:
			fmt.Printf("\n%d.) %s [%c]\n%s\n", id, title, checkBoxStatus, description)
		case false:
			fmt.Printf("\n%d.) %s [%c]\n%s\n", id, title, checkBoxStatus, description)
		}

		for _, char := range line {
			fmt.Printf("%c", char)
		}

		isFirstColumn = !isFirstColumn
	}

	fmt.Println()

	return nil // Successfully printed table
}
