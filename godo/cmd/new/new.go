/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package new

import (
	"fmt"

	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/cmd/new/subcmds"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new task or task group",
	Long: `

This command creates either a task or task group.
If no task group is explicitly given, the task 
will be added to a default task group

Example:
	godo new task "Add docstrings" // Creates a task in the default task group
	godo new task "Create tests for this file" "tests" // Creates a task in the tests task group
	
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
	newCmd.AddCommand(subcmds.GroupCmd)
	newCmd.AddCommand(subcmds.TaskCmd)
}
