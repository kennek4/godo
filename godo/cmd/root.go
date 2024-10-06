/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	CurrentWorkingDirectory string

	// The directory location of .godo
	GodoDir string

	// This path ends in a godo.db
	GodoDbPath string
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "godo",
	Short: "A todo list organizer to keep track of tasks in your current working directory",
	Long:  `A CLI todo list organizer made in Go to keep track of tasks you need to do in your development environment.`,
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
	currWd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	CurrentWorkingDirectory = currWd
	GodoDir = filepath.Join(currWd, ".godo")
	GodoDbPath = filepath.Join(GodoDir, "godo.db")

	viper.SetEnvPrefix("godo")
	viper.BindEnv("group")
}
