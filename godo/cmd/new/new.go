/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package new

import (
	"fmt"

	"github.com/kennek4/godo/cmd"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var NewCmd = &cobra.Command{
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")
	},
}

func init() {
	cmd.RootCmd.AddCommand(NewCmd)
}
