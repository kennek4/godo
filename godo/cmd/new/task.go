/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package new

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kennek4/godo/cmd"
	"github.com/spf13/cobra"
)

var (
	taskName    string
	taskDetails string
)

// taskCmd represents the task command and is a subcommand of new
var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Creates a task",
	Long: `

Creates a task with a name and description 

You do not need quotation marks if your arguments
do not have any spaces

`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		taskName = args[0]
		taskDetails = args[1]
		createTask()
	},
}

type Task struct {
	TaskID      int
	TaskName    string
	TaskDetails string
	IsCompleted bool
}

func createID() (taskID int, err error) {
	// Count how many tasks already exist
	pattern := filepath.Join(cmd.GodoPath+".godo", "*.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return 0, err
	}

	return len(files) + 1, nil
}

func createTask() error {
	taskItem := Task{}

	taskID, err := createID()
	if err != nil {
		return err
	}

	// Set fields
	taskItem.TaskID = taskID
	taskItem.TaskName = taskName
	taskItem.TaskDetails = taskDetails
	taskItem.IsCompleted = false

	task, err := json.MarshalIndent(taskItem, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fileName := fmt.Sprintf(".godo/Task_%d.json", taskID)
	err = os.WriteFile(fileName, task, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	taskCmd.PersistentFlags().StringVarP(&taskName, "name", "n", "", "name of the task")
	taskCmd.PersistentFlags().StringVarP(&taskDetails, "details", "d", "", "details of the task")
	newCmd.AddCommand(taskCmd)
}
