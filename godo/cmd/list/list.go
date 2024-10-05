/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package list

import (
	"fmt"

	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists godo tasks",
	Long: `This command lists godo tasks, by default prints out all tasks.
An optional argument can be provided to sort the category of tasks to show.

Example:

The following command will show all tasks under the "Code" category.

> godo list Code
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskTable := args[0]
		listTasks(taskTable)
	},
}

func listTasks(taskTable string) error {

	db, err := util.GetDB(&cmd.GodoDir)
	if err != nil {
		return err
	}

	defer db.Close()

	prep := fmt.Sprintf("INSERT INTO %s (title, description) VALUES (?, ?)", taskTable)
	statement, err := db.Prepare(prep)
	if err != nil {
		return err
	}

	statement.Exec("Make Cookies", "Make by tonight!")

	query := fmt.Sprintf("SELECT id, title, description FROM %s", taskTable)
	rows, _ := db.Query(query)

	var id int
	var title string
	var description string
	for rows.Next() {
		rows.Scan(&id, &title, &description)
		fmt.Printf("[%d] %s\n%s\n", id, title, description)
	}

	return nil
}

func init() {
	cmd.RootCmd.AddCommand(listCmd)
}
