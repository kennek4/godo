/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kennek4/godo/internal/util/configs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	CurrentWorkingDirectory string

	// The directory location of .godo
	GodoDir string

	// This path ends in a godo.db
	GodoDbPath string

	CurrentGroup string
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

func setDefaults() {
	viper.SetDefault("group", "tasks")
}

func init() {

	currWd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	CurrentWorkingDirectory = currWd
	GodoDir = filepath.Join(currWd, ".godo")
	GodoDbPath = filepath.Join(GodoDir, "godo.db")

	cobra.OnInitialize(initConfig)
	setDefaults()
}

func initConfig() {
	if cfgFile != "" {

		cfgData, err := configs.GetConfig(GodoDir)
		if err != nil {
			log.Fatal(err)
		}

		CurrentGroup = cfgData.CurrentGroup
		fmt.Println(CurrentGroup)

	} else {
	}
}
