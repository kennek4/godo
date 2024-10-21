/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package list

import (
	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/configs"
	"github.com/kennek4/godo/internal/util/gddb"
	"github.com/kennek4/godo/internal/util/gdmisc"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists godo tasks",
	Long: `This command lists godo tasks, by default prints out all tasks.
An optional argument can be provided to sort the group of tasks to show.

Example:

The following command will show all tasks under the "Code" category.

> godo list Code
`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(command *cobra.Command, args []string) error {
		taskTable := configs.GetCurrentGroup(cmd.GodoDir)

		if len(args) > 0 {
			taskTable = args[0]
		}
		tasks, err := gddb.ListTasksInTable(&taskTable, &cmd.GodoDir)
		if err != nil {
			return err
		}
		gdmisc.DisplayTasks(tasks, &taskTable)
		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(listCmd)
}
