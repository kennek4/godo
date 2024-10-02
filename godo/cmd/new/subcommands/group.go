/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package new

import (
	"fmt"

	"github.com/kennek4/godo/cmd/new"
	"github.com/spf13/cobra"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("group called")
	},
}

func init() {
	new.NewCmd.AddCommand(groupCmd)
}
