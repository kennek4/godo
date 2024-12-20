/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package new

import (
	"github.com/kennek4/genv"
	"github.com/kennek4/godo/internal/util/gddb"
	"github.com/spf13/cobra"
)

var (
	taskTitle       string
	taskDescription string
)

// taskCmd represents the task command and is a subcommand of new
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Creates a task",
	Long: `

Creates a task with a name and description 

You do not need quotation marks if your arguments
do not have any spaces.

`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskTitle = args[0]
		taskDescription = args[1]
		createTask(&taskTitle, &taskDescription)
	},
}

func createTask(title *string, desc *string) error {
	group := genv.GetVar("CurrentGroup")
	err := gddb.CreateTask(*title, *desc, &group)
	if err != nil {
		return err
	}

	return nil // Task created successfully
}

func init() {
	taskCmd.PersistentFlags().StringVarP(&taskTitle, "title", "t", "", "name of the task")
	taskCmd.PersistentFlags().StringVarP(&taskDescription, "description", "d", "", "details of the task")
	newCmd.AddCommand(taskCmd)
}
