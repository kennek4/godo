/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package new

import (
	"fmt"

	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/configs"
	"github.com/kennek4/godo/internal/util/dbdriver"
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
		taskDescription = "\t" + args[1]
		createTask(&taskTitle, &taskDescription)
	},
}

func createTask(title *string, description *string) error {

	table := configs.GetValueFromKey("group")
	fmt.Printf("table: %s\n", table)
	err := dbdriver.InsertTaskInDB(title, description, &table, &cmd.GodoDir)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	taskCmd.PersistentFlags().StringVarP(&taskTitle, "title", "t", "", "name of the task")
	taskCmd.PersistentFlags().StringVarP(&taskDescription, "description", "d", "", "details of the task")
	newCmd.AddCommand(taskCmd)
}
