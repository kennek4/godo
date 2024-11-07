package gdmisc

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/kennek4/godo/internal/util/gddb"
)


func ClearConsole() {

	var cmd *exec.Cmd
	goos := runtime.GOOS

	switch goos {
	case "linux":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func DisplayGroups(groups []gddb.Group) error {

	ClearConsole()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetTitle(text.FormatUpper.Apply("List of Groups"))
	t.AppendHeader(table.Row{"Group Name", "Task Count"})

	for _, group := range groups {
		t.AppendRow(table.Row{group.Name, group.TaskCount})
		t.AppendSeparator()
	}

	t.SetStyle(table.StyleBold)
	t.Render()

	return nil
}

func DisplayTasks(tasks []gddb.Task, tableName *string) error {

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

// Prompts the user if they would like to proceed by asking Y/n
func PromptUser(prompt *string) (userResponse bool, err error) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s", *prompt)
	fmt.Print("Would you like to proceed? [Y/n] ")

	char, _, err := reader.ReadRune()
	if err != nil { // Rune could not be encoded to UTF-8
		return false, fmt.Errorf("rune could not be encoded to UTF-8")
	}

	switch char {
	case 'Y':
	case 'y':
		return true, nil

	case 'N':
	case 'n':
		return false, nil
	}

	return false, err
}
