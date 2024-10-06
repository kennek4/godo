/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package init

import (
	"fmt"
	"os"

	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/consolehelper"
	"github.com/kennek4/godo/internal/util/dbdriver"
	"github.com/kennek4/godo/internal/util/filehelper"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a godo instance in the current directory",
	Long:  "This command will initialize a godo into the current directory in a hidden folder called .godo",
	Args:  cobra.ExactArgs(0),
	RunE: func(command *cobra.Command, args []string) error {

		// Get any flags if they exist
		confirmFlag, err := command.Flags().GetBool("confirm")
		if err != nil {
			return err
		}

		var choice bool

		switch confirmFlag {
		case true:
			err = startInit(true)
		default:
			prompt := "godo requires a sqlite database\n"
			choice, err = consolehelper.PromptUser(&prompt)
			if err != nil {
				return err
			}

			err = startInit(choice)
		}

		if err != nil {
			return err
		}

		return nil
	},
}

func startInit(choice bool) error {
	var err error

	switch {
	case choice: // User accepted prompt
		err = initGodo(&cmd.GodoDir)
	case !choice: // User denied prompt
		return fmt.Errorf("godo requires a sqlite database to work, please try again")
	}

	if err != nil {
		return err
	}

	return nil
}

func initGodo(godoDir *string) error {

	// Define where .godo will be placed

	if _, err := os.Stat(*godoDir); err == nil {
		// go-do is already initialized
		return fmt.Errorf("godo is already initialized in this directory")
	}

	// Create new .godo directory
	err := os.Mkdir(*godoDir, 0777)
	if err != nil && os.IsNotExist(err) {
		return err
	}

	err = filehelper.MakeDirHidden(godoDir)
	if err != nil {
		return err
	}

	// Initialize DB in .godo
	err = dbdriver.InitDB(godoDir)
	if err != nil {
		return err
	}

	// Add files to .gitignore (if present)
	return nil
}

func init() {
	cmd.RootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("confirm", "c", false, "Used to accept the prompt without it being shown")
}
