/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package task

import (
	"fmt"
	"log"

	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/configs"
	"github.com/kennek4/godo/internal/util/gddb"
	"github.com/kennek4/godo/internal/util/gdmisc"
	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Used for task manipulations",
	Long: `This task is the base command for task manipulations such as viewing a certain task,
completing a certain task, editing a task, etc.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := displayMenu()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func getChoice() (choice int, err error) {
	fmt.Scanf("%d", &choice)
	return choice, nil
}

func displayMenu() error {

	gdmisc.ClearConsole()

	fmt.Println(`
EDIT MENU
PLEASE SELECT ONE OF THE FOLLOWING

[1] Change task title
[2] Change task description
[3] Complete task
	`)

	var choice int

	if input, err := getChoice(); err == nil {
		choice = input
	} else {
		return err
	}

	switch choice {
	case 1:
		changeTitle()
	case 2:
		// changeDescription()
	case 3:
		// completeTask()
	default:
		fmt.Println("Invalid input. Please enter a value from 1-3.")
	}

	return nil
}

func changeTitle() {

	var newTitle string

	table := configs.GetCurrentGroup(cmd.GodoDir)
	tasks, err := gddb.GetTasksFromDB(&table, &cmd.GodoDir)
	if err != nil {
		log.Fatal(err)
	}

	gdmisc.DisplayTasks(tasks, &table)

	var choice int
	if input, err := getChoice(); err == nil {
		choice = input
	} else {
		log.Fatal(err)
	}

	fmt.Printf("choice: %v\n", choice)

	err = gddb.EditTaskTitle(choice, &newTitle, &table, &cmd.GodoDir)
	if err != nil {
		log.Fatal(err)
	}

}

func init() {
	cmd.RootCmd.AddCommand(taskCmd)
}
