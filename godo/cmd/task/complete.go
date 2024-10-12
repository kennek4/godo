/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package task

import (
	"fmt"
	"strconv"

	"github.com/kennek4/godo/cmd"
	"github.com/kennek4/godo/internal/util/configs"
	"github.com/kennek4/godo/internal/util/dbdriver"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return fmt.Errorf("one argument is required, a Task ID")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		taskId, err := isInt(args[0])
		if err != nil {
			return err
		}

		err = processCompleteCmd(&taskId)
		if err != nil {
			return err
		}

		return nil
	},
}

func processCompleteCmd(taskId *int) error {
	currentGroup := configs.GetCurrentGroup(cmd.GodoDir)

	// Change completion status on task
	err := dbdriver.CompleteTask(taskId, &currentGroup, &cmd.GodoDir)
	if err != nil {
		return err
	}

	// List all tasks
	err = dbdriver.ListTasksInTable(&currentGroup, &cmd.GodoDir)
	if err != nil {
		return err
	}
	
	return nil
}

func isInt(s string) (num int, err error) {
	num, err = strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("the given argument is not a valid Task ID (int)")
	}

	return num, nil
}

func init() {
	TaskCmd.AddCommand(completeCmd)
}
