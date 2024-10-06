/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package new

import (
	"fmt"

	"github.com/kennek4/godo/cmd"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new task",
	Long: `

This command creates either a task or task group.
If no task group is explicitly given, the task 
will be added to a default task group

Example:

This will create a new task with the title 
> godo new task "Add docstrings"
	
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MaximumNArgs(1); err != nil {
			return fmt.Errorf("too many arguments given, provide only one argument when using this command")
		}

		if err := cobra.MinimumNArgs(1); err != nil {
			return fmt.Errorf("no arguments given, please enter one when using this command")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(newCmd)
}
