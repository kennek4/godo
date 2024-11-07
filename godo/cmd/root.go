/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/kennek4/genv"
	"github.com/kennek4/godo/internal/util/gddb"
	"github.com/spf13/cobra"
)

const (
	AppName string = "Godo"
	EnvFile string = ".GODO.env"
)

var (
	EnvFilePath string
	GodoDir     string
)

var ()

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "godo",
	Short: "A todo list organizer to keep track of tasks in your current working directory",
	Long:  `A CLI todo list organizer made in Go to keep track of tasks you need to do in your development environment.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		genv.Save()

		if gddb.GodoDb != nil {
			gddb.Close(gddb.GodoDb)
		}

		os.Exit(1)
	}
}
func init() {

	home, _ := os.UserHomeDir()
	titleAppName := strings.ToTitle(AppName)

	GodoDir = filepath.Join(home, "."+titleAppName)
	EnvFilePath = filepath.Join(GodoDir, EnvFile)

	// Check if Godo directory exists
	_, err := os.Stat(EnvFilePath)
	switch os.IsNotExist(err) {
	case true:
		err = genv.Init(AppName)
		if err != nil {
			log.Fatalf("Failed to initialize godo, %s", err)
		}
		genv.CreateStringVar("CurrentGroup", "")

	case false:
		err = genv.Load(AppName)
		if err != nil {
			log.Fatalf("Failed to load godo config file, %s", err)
		}

		err = gddb.Load(GodoDir)
		if err != nil {
			log.Fatalf("Failed to load godo.db, %s", err)
		}

	default:
		log.Fatal("Something went catastrophically wrong during root.init()")
	}
}
