/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package group

import (
	"fmt"

	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/configs"
	"github.com/kennek4/godo/internal/util/console"
	"github.com/kennek4/godo/internal/util/dbdriver"
	"github.com/spf13/cobra"
)

var tables []string

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "This command is used for godo group functions",
	Long: `By default this command is used to change the current group for godo.
The following command would change the godo group to one named "Code"
> godo group Code`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(command *cobra.Command, args []string) error {

		tables = dbdriver.ListTablesInDB(&cmd.GodoDir)
		groupName := args[0]

		switch doesGroupExist(groupName) {
		case false:
			prompt := fmt.Sprintf("The table, %s, does not exist in the godo db.\nWould you like to create a group names %s?\n", groupName, groupName)
			choice, err := console.PromptUser(&prompt)
			if err != nil {
				return err
			}

			if choice { // Create new group
				err = dbdriver.CreateTableInDB(groupName, &cmd.GodoDir)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("the given group was not found, please try again")
			}

		case true:
			configs.SetCurrentGroup(groupName, cmd.GodoDir) // Set new CurrentGroup
		}

		return nil
	},
}

func doesGroupExist(groupName string) bool {
	for _, value := range tables {
		if value == groupName {
			return true
		}
	}

	return false
}

func init() {
	cmd.RootCmd.AddCommand(groupCmd)
}
