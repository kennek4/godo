/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package delete

import (
	"github.com/kennek4/genv"
	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/gddb"
	"github.com/spf13/cobra"
)

var (
	arguments []string
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

		titles, _ := cmd.Flags().GetStringArray("title")
		ids, _ := cmd.Flags().GetStringArray("id")

		// Only one of these will have values
		if len(titles) > 0 {
			err := deleteTask(gddb.Title, titles)
			if err != nil {
				return err
			}
		}

		if len(ids) > 0 {
			err := deleteTask(gddb.Id, ids)
			if err != nil {
				return err
			}
		}

		return nil
	},
}

func deleteTask(deleteType gddb.DeleteType, taskProperties []string) error {

	group := genv.GetVar("CurrentGroup")
	err := gddb.Delete(deleteType, group, taskProperties)
	if err != nil {
		return err
	}

	return nil // Task(s) deleted
}

func init() {
	deleteCmd.Flags().StringArrayVarP(&arguments, "title", "t", arguments, "The title of the task the user wishes to delete")
	deleteCmd.Flags().StringArrayVarP(&arguments, "id", "i", arguments, "The ID of the task the user wishes to delete")

	// Either is required but both flags can not be set
	deleteCmd.MarkFlagsOneRequired("title", "id")
	deleteCmd.MarkFlagsMutuallyExclusive("title", "id")

	cmd.RootCmd.AddCommand(deleteCmd)
}
