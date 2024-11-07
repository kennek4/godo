/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package group

import (
	"github.com/kennek4/godo/internal/util/gddb"
	"github.com/kennek4/godo/internal/util/gdmisc"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A subcommand of group that lists all groups",
	Long: `This command will list all groups that have been made in Godo. 
it will also show how many tasks are in the current group.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := listGroups()
		if err != nil {
			return err
		}

		return nil
	},
}

func listGroups() error {
	groups, err := gddb.GetGroup()
	if err != nil {
		return err
	}

	gdmisc.DisplayGroups(groups)

	return nil // Groups successfully displayed
}

func init() {
	groupCmd.AddCommand(listCmd)
}
