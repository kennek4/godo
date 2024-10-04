/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package init

import (
	"fmt"
	"os"
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
		if cmd.IsInitalized {
			return fmt.Errorf("a godo instance is already intialized in %s", cmd.GodoPath)
		}

		prompt := fmt.Sprintf("godo will initialize in: %s\ngodo will create any necessary directories\n", cmd.GodoPath)
		choice, err := util.PromptUser(prompt)
		if err != nil {
			return err
		}

		if choice {
			// User accepted prompt
			err = initGodo(cmd.GodoPath)
			if err != nil {
				return err
			}
		} else {
			// User denied prompt
			return fmt.Errorf("godo requires a directory to work, please try again")
		}

		return nil
	},
}

func init() {
	cmd.RootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().StringVarP(&cmd.GodoPath, "directory", "d", "./", "the directory in which godo will work in")
	// Bind Flags to Viper
	viper.BindPFlag("directory", initCmd.PersistentFlags().Lookup("directory"))
}

func initGodo(dirPath string) error {

	// Define where .godo will be placed
	godoPath := fmt.Sprintf("%s/.godo", dirPath)

	if _, err := os.Stat(godoPath); err == nil {
		// go-do is already initialized
		return fmt.Errorf("godo is already initialized in this directory")
	}

	// Create new .godo directory
	err := os.MkdirAll(godoPath, 0777)
	if err != nil && os.IsNotExist(err) {
		return err
	}

	// Set .godo to HIDDEN
	newDirPtr, err := syscall.UTF16PtrFromString(godoPath)
	if err != nil {
		return err
	}

	err = syscall.SetFileAttributes(newDirPtr, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		return err
	}

	cmd.IsInitalized = true

	// Add files to .gitignore (if present)
	return nil
}
