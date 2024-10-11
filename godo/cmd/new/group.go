/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package new

import (
	"fmt"

	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/dbdriver"
	"github.com/spf13/cobra"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Creates a new task group",
	Long: `This command creates a new task group for task organization.
The following command would create a new task group called "Homework"
> godo new group Homework
`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		groupName := args[0]
		if groupName == "" {
			return fmt.Errorf("can't create a new group with an empty name, please try again")
		}

		if groupName != "" {
			err := createGroup(&groupName)
			if err != nil {
				return err
			}

			return nil // Group was successfully created
		}

		return fmt.Errorf("failed to create a new table")
	},
}

func createGroup(groupName *string) error {

	if groupName == nil {
		err := fmt.Errorf("in createGroup, groupName was supplied a nil string pointer")
		return err
	}

	dbdriver.CreateTableInDB(*groupName, &cmd.GodoDir)
	return nil
}

func init() {
	groupCmd.PersistentFlags().StringP("groupName", "n", "", "The name of the group that the user wishes to create")
	newCmd.AddCommand(groupCmd)
}
