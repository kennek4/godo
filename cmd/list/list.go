/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package list

import (
	"fmt"

	"github.com/kennek4/genv"
	"github.com/kennek4/godo/cmd"
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

		groupFlag, err := command.Flags().GetBool("group")
		if err != nil {
			return err
		}

		switch groupFlag {
		case true: // Print Groups
			groups, err := gddb.GetGroups()
			if err != nil {
				return fmt.Errorf("failed to get groups from Godo, %s", err)
			}

			gdmisc.DisplayGroups(groups)

		case false: // Print Tasks
			taskGroup := genv.GetVar("CurrentGroup")
			tasks, err := gddb.GetTasks(taskGroup)
			if err != nil {
				return fmt.Errorf("failed to get tasks from Godo, %s", err)
			}

			gdmisc.DisplayTasks(tasks, &taskGroup)
		}

		return nil
	},
}

func init() {
	listCmd.Flags().BoolP("group", "g", false, "Used to list all Godo groups")
	cmd.RootCmd.AddCommand(listCmd)
}
