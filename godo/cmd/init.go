/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a godo instance in the current directory",
	Long:  "This command will initialize a godo into the current directory in a hidden folder called .godo",
	RunE: func(cmd *cobra.Command, args []string) error {

		err := initGodo()
		if err != nil || isInitalized {
			return fmt.Errorf("a godo instance is already intialized at: %s", todoDir)
		} else {
			isInitalized = true
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.PersistentFlags().StringVarP(&todoDir, "directory", "d", "./", `the directory in which godo will work in. 
																by default this will use the current working directory`)

	// Bind Flags to Viper
	viper.BindPFlag("directory", initCmd.PersistentFlags().Lookup("directory"))
}
