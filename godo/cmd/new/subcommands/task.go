/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package new

import (
	"fmt"

	"github.com/kennek4/godo/cmd/new"
	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("task called")
	},
}

func init() {
	new.NewCmd.AddCommand(taskCmd)

}
