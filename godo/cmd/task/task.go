/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package task

import (
	"github.com/kennek4/godo/cmd"
	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Used for task manipulations",
	Long: `This task is the base command for task manipulations such as viewing a certain task,
completing a certain task, editing a task, etc.
This command does nothing without a verb command following it.
`,
	Args: cobra.ExactArgs(0),
}

func init() {
	cmd.RootCmd.AddCommand(TaskCmd)
}
