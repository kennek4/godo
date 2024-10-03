/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package init

import (
	"fmt"
	"os"
	"syscall"

	"github.com/kennek4/godo/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// The directory location of godo
	todoDir string

	// The file name
	todoFile string

	// Boolean to see if godo is already initialized
	isInitalized bool
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
	cmd.RootCmd.AddCommand(initCmd)

	initCmd.PersistentFlags().StringVarP(&todoDir, "directory", "d", "./", `the directory in which godo will work in. 
																by default this will use the current working directory`)
	// Bind Flags to Viper
	viper.BindPFlag("directory", initCmd.PersistentFlags().Lookup("directory"))
}

func initGodo() error {
	todoFilePath := fmt.Sprintf("%s.godo/todos.json", todoDir)
	_, err := os.Stat(todoFilePath)
	if err == nil {
		// go-do is already initialized
		return fmt.Errorf("godo is already initialized in this directory")
	}

	// Create files
	err = createFiles()
	if err != nil {
		return err
	}

	// Add files to .gitignore (if present)
	return nil
}

func createFiles() error {
	todoFile = viper.GetString("todoFile") // "todos.json"

	newDirPath := fmt.Sprintf("%s.godo", todoDir)

	// Create new .godo directory
	err := os.Mkdir(newDirPath, 0777)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	newDirPtr, err := syscall.UTF16PtrFromString(newDirPath)
	if err != nil {
		return err
	}

	// Set .godo to HIDDEN
	err = syscall.SetFileAttributes(newDirPtr, syscall.FILE_ATTRIBUTE_HIDDEN)
	if err != nil {
		return err
	}

	var newFilePath string = fmt.Sprintf("%s%s", todoDir, todoFile)
	var file *os.File

	file, err = os.Create(newFilePath)
	if err != nil {
		return err // File could not be created
	}

	file.Close()

	return nil
}
