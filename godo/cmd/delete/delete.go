/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/dbdriver"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Used to delete tasks or groups",
	Long: `This task is used to delete tasks or groups from godo
	
When provided with no flags, the default argument will be for the index of the task
that the user wishes to delete.
The following command will delete a task with an ID of 3
> godo delete 3

To delete a group, the -g flag should be provided then name of the group that you wish to delete.
The following example will delete a group called "Backend"
> godo delete -g Backend

To delete a task by its task, please use the -t flag with the title of the task.
The following example will delete a task with the title "Review Code"
> godo delete -t "Review Code"
`,
	RunE: func(cmd *cobra.Command, args []string) error {

		var thingToDelete interface{}

		// groupFlag, _ := cmd.Flags().GetString("group")
		titleFlag, _ := cmd.Flags().GetString("title")

		if titleFlag != "" {
			thingToDelete = titleFlag
			err := deleteTask(1, &thingToDelete)
			if err != nil {
				return err
			}
		} else {
			err := deleteTask(2, &thingToDelete)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func deleteTask(quertyType dbdriver.DeleteType, thingToDelete *interface{}) error {
	err := dbdriver.DeleteTaskInDB(quertyType, "cheese", thingToDelete, &cmd.GodoDir)
	if err != nil {
		return err
	}

	return nil // Task deleted
}

func init() {
	// deleteCmd.Flags().StringP("group", "g", "", "The group the user wishes to delete")
	deleteCmd.Flags().StringP("title", "t", "", "The title of the task the user wishes to delete")
	cmd.RootCmd.AddCommand(deleteCmd)
}
