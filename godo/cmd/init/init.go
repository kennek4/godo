/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package init

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a godo instance in the current directory",
	Long:  "This command will initialize a godo into the current directory in a hidden folder called .godo",
	Args:  cobra.ExactArgs(0),
	RunE: func(command *cobra.Command, args []string) error {

		prompt := "godo requires a sqlite database\n"
		choice, err := util.PromptUser(prompt)
		if err != nil {
			return err
		}

		switch {
		case choice: // User accepted prompt
			err = initGodo(cmd.GodoDir)
			if err != nil {
				return err
			}
		case !choice: // User denied prompt
			return fmt.Errorf("godo requires a directory to work, please try again")
		}

		return nil
	},
}

func initGodo(currWd string) error {

	// Define where .godo will be placed
	godoDir := filepath.Join(currWd, ".godo")

	if _, err := os.Stat(godoDir); err == nil {
		// go-do is already initialized
		return fmt.Errorf("godo is already initialized in this directory")
	}

	// Create new .godo directory
	err := os.MkdirAll(godoDir, 0777)
	if err != nil && os.IsNotExist(err) {
		return err
	}

	newDirPtr, err := syscall.UTF16PtrFromString(godoDir)
	if err != nil {
		return err
	}

	// Set directory to hidden
	err = syscall.SetFileAttributes(newDirPtr, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		return err
	}

	// Initialize DB in .godo
	err = util.InitDB(cmd.CurrentWorkingDirectory)
	if err != nil {
		return err
	}

	// Add files to .gitignore (if present)
	return nil
}

func init() {
	cmd.RootCmd.AddCommand(initCmd)

	// Bind Flags to Viper
	viper.BindPFlag("directory", initCmd.PersistentFlags().Lookup("directory"))
}
