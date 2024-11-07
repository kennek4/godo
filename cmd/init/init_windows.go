/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package init

import (
	"fmt"

	"github.com/kennek4/genv"
	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/gddb"
	"github.com/kennek4/godo/internal/util/gdmisc"
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

		defaultGroup, err := command.Flags().GetString("group")
		if err != nil {
			return nil
		}

		var choice bool

		switch confirmFlag {
		case true:
			err = startInit(defaultGroup, true)
		default:
			prompt := "godo requires a sqlite database\n"
			choice, err = gdmisc.PromptUser(&prompt)
			if err != nil {
				return err
			}

			err = startInit(defaultGroup, choice)
		}

		if err != nil {
			return err
		}

		return nil
	},
}

func startInit(defaultGroup string, choice bool) error {
	var err error

	switch {
	case choice: // User accepted prompt
		err = initGodo(defaultGroup, &cmd.GodoDir)
	case !choice: // User denied prompt
		return fmt.Errorf("godo requires a sqlite database to work, please try again")
	}

	if err != nil {
		return err
	}

	return nil
}

func initGodo(defaultTable string, godoDir *string) error {

	if godoDir == nil {
		return fmt.Errorf("in initGodo, godoDir was passed a nil string pointer")
	}

	// Initialize DB in .godo
	err := gddb.Load(*godoDir)
	if err != nil {
		return err
	}

	genv.CreateStringVar("CurrentGroup", defaultTable)

	return nil
}

func init() {
	cmd.RootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolP("confirm", "c", false, "Used to accept the prompt without it being shown")
	initCmd.Flags().StringP("group", "g", "tasks", "Used to set the default SQLite database table")
}
