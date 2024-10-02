/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"syscall"

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

func initGodo() error {
	todoFilePath := fmt.Sprintf("%s.godo/todos.json", todoDir)
	_, err := os.Stat(todoFilePath)
	if err == nil {
		// go-do is already initialized
		return fmt.Errorf("godo is already initialized in this directory")
	} else {
		// go-do is not initialized

		todoFile = "todos.json"

		newDirPath := fmt.Sprintf("%s.godo", todoDir)

		// Create new .godo directory
		err := os.Mkdir(newDirPath, 0777)
		if err != nil && !os.IsExist(err) {
			panic(err)
		}

		newDirPtr, err := syscall.UTF16PtrFromString(newDirPath)
		if err != nil {
			return err
		}

		// Set .godo to HIDDEN
		err = syscall.SetFileAttributes(newDirPtr, syscall.FILE_ATTRIBUTE_HIDDEN)
		if err != nil {
			return err
		}

		_, err = os.Create(todoFile)
		if err != nil {
			return err // File could not be created
		}

		return nil
	}
}
