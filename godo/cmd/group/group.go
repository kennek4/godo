/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package group

import (
	"fmt"

	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/configs"
	"github.com/kennek4/godo/internal/util/gddb"
	"github.com/spf13/cobra"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "This command is used for godo group functions",
	Long: `By default this command is used to change the current group for godo.
The following command would change the godo group to one named "Code"
> godo group Code`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(command *cobra.Command, args []string) error {

		groups, err := gddb.GetGroupsFromDB(&cmd.GodoDir)
		if err != nil {
			return err
		}
		groupName := args[0]

		if !doesGroupExist(groups, groupName) { // Given argument is not a table in the DB
			err := fmt.Errorf("the table, %s, does not exist in the godo db, please check your spelling or try again", groupName)
			return err
		} else { // Given argument exists
			configs.SetCurrentGroup(groupName, cmd.GodoDir) // Set new CurrentGroup
		}

		return nil
	},
}

func doesGroupExist(groups []gddb.Group, groupName string) bool {
	for _, group := range groups {
		if group.Name == groupName {
			return true
		}
	}

	return false
}

func init() {
	cmd.RootCmd.AddCommand(groupCmd)
}
