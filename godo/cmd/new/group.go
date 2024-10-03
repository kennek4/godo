/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package new

import (
	"fmt"

	"github.com/spf13/cobra"
)

var groupName string

// groupCmd represents the group command and is a subcommand of new
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Creates a task group",
	Long: `

Creates a task group that tasks can be organized into.
This command only takes one, and only one argument
that represents the name of the task group.

The argument provided does not need to be surrounded 
in quotation marks (such as "" or '').

	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("group called with arg: %s", args[0])
	},
}

// func createTaskGroup() error {

// 	// var currDirPath string = viper.GetString("dirPath")
// 	// dirPath, err := os.Stat(fmt.Sprintf("%s.godo", currDirPath))

// }

func init() {
	groupCmd.PersistentFlags().StringVarP(&groupName, "name", "n", "", "name of the task group")
	newCmd.AddCommand(groupCmd)

}
