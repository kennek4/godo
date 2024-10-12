/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package list

import (
	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/configs"
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
	Args: cobra.MaximumNArgs(1),
	Run: func(command *cobra.Command, args []string) {
		if len(args) > 0 {
			listTasks(&args[0])
		} else {
			currentGroup := configs.GetCurrentGroup(cmd.GodoDir)
			listTasks(&currentGroup)
		}
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
