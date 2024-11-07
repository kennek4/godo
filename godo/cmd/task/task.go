/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package task

import (
	"fmt"
	"log"

	"github.com/kennek4/godo/cmd"
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
	return nil
}

func init() {
	cmd.RootCmd.AddCommand(taskCmd)
}
