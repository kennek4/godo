package consolehelper

import (
	"os"
	"os/exec"
)

func ClearConsole() {
	clear := exec.Command("cmd", "/c", "cls")
	clear.Stdout = os.Stdout
	clear.Run()
}