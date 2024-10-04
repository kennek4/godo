/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package task

import (
	"fmt"

	"github.com/kennek4/godo/cmd"
	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Used for task manipulations",
	Long: `This task is the base command for task manipulations such as viewing a certain task,
completing a certain task, editing a task, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(taskCmd)
}
