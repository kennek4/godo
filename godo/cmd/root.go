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
	GodoPath string

	// Boolean to see if godo is already initialized
	IsInitalized bool
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "godo",
	Short: "A todo list organizer to keep track of tasks in your current working directory",
	Long:  `A CLI todo list organizer made in Go to keep track of tasks you need to do in your development environment.`,
}

func printTasks() {

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	currWd, _ := os.Getwd()
	viper.SetDefault("godoPath", currWd)
}
