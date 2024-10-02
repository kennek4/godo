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

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a godo instance in the current directory",
	Long:  "This command will initialize a godo into the current directory in a hidden folder called .godo",
	RunE: func(cmd *cobra.Command, args []string) error {

		err := initGodo()
		if err != nil || isInitalized {
			return fmt.Errorf("a godo instance is already intialized at: %s", todoDir)
		} else {
			isInitalized = true
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.PersistentFlags().StringVarP(&todoDir, "directory", "d", "./", `the directory in which godo will work in. 
																by default this will use the current working directory`)

	// Bind Flags to Viper
	viper.BindPFlag("directory", initCmd.PersistentFlags().Lookup("directory"))
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
