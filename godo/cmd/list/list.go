/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/dbdriver"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists godo tasks",
	Long: `This command lists godo tasks, by default prints out all tasks.
An optional argument can be provided to sort the category of tasks to show.

Example:

The following command will show all tasks under the "Code" category.

> godo list Code
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskTable := args[0]
		listTasks(&taskTable)
	},
}

func listTasks(table *string) error {

	err := dbdriver.ListTasksInTable(table, &cmd.GodoDir)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	cmd.RootCmd.AddCommand(listCmd)
}
