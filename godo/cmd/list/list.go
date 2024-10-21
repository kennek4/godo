/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package list

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/configs"
	"github.com/kennek4/godo/internal/util/gddb"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists godo tasks",
	Long: `This command lists godo tasks, by default prints out all tasks.
An optional argument can be provided to sort the group of tasks to show.

Example:

The following command will show all tasks under the "Code" category.

> godo list Code
`,
	Args: cobra.MaximumNArgs(1),
	Run: func(command *cobra.Command, args []string) {
		taskTable := configs.GetCurrentGroup(cmd.GodoDir)

		if len(args) > 0 {
			taskTable = args[0]
		}

		listTasks(&taskTable)
	},
}

func listTasks(tableName *string) error {

	tasks, err := gddb.ListTasksInTable(tableName, &cmd.GodoDir)
	if err != nil {
		return err
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle(text.FormatTitle.Apply(*tableName))
	t.AppendHeader(table.Row{"Task #", "Title", "Description", "Complete"})

	var completeIcon rune
	var completeness string
	for _, task := range tasks {

		switch task.IsComplete {
		case true:
			completeIcon = '✔'
			completeness = text.FgHiGreen.Sprintf("%c", completeIcon)
		case false:
			completeIcon = '✘'
			completeness = text.FgHiRed.Sprintf("%c", completeIcon)
		}

		title := text.WrapSoft(task.Title, 30)
		description := text.WrapSoft(task.Description, 30)

		t.AppendRow(table.Row{
			task.Id,
			title,
			description,
			completeness,
		})

		t.AppendSeparator()
	}

	t.SetStyle(table.StyleBold)
	t.Render()

	return nil
}

func init() {
	cmd.RootCmd.AddCommand(listCmd)
}
