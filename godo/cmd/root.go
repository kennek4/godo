/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// The directory location of godo
	todoDir string

	// The file name
	todoFile string

	// Boolean to see if godo is already initialized
	isInitalized bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "godo",
	Short: "A todo list organizer to keep track of tasks in your current working directory",
	Long:  `A CLI todo list organizer made in Go to keep track of tasks you need to do in your development environment.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	viper.SetDefault("todoFile", "todos.json")

	// Adding subcommands
	rootCmd.AddCommand(initCmd)
}
