/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package new

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kennek4/genv"
	"github.com/kennek4/godo/internal/util/gddb"
	"github.com/spf13/cobra"
)

var group string

// taskCmd represents the task command and is a subcommand of new
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Creates a task",
	Long: `

Creates a task with a name and description 

You do not need quotation marks if your arguments
do not have any spaces.

`,
	Args: cobra.MaximumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		titleFlag, err := cmd.Flags().GetString("title")
		if err != nil {
			return err
		}

		descFlag, err := cmd.Flags().GetString("desc")
		if err != nil {
			return err
		}

		flagsProvided := titleFlag == "" && descFlag == ""
		group = genv.GetVar("CurrentGroup")

		switch flagsProvided {
		case true:
			err = createTask()
			if err != nil {
				return err
			}
		case false:
			err = gddb.CreateTask(titleFlag, descFlag, &group)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func createTask() error {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Task Title: ")
	taskTitle, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Print("Description: ")
	taskDesc, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	err = gddb.CreateTask(taskTitle, taskDesc, &group)
	if err != nil {
		return err
	}

	return nil // Task created successfully
}

func init() {
	taskCmd.Flags().StringP("title", "t", "", "name of the task")
	taskCmd.Flags().StringP("desc", "d", "", "details of the task")

	taskCmd.MarkFlagsRequiredTogether("title", "desc")

	newCmd.AddCommand(taskCmd)
}
