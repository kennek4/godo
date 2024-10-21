/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package main

import (
	"github.com/kennek4/godo/cmd"
	_ "github.com/kennek4/godo/cmd/delete"
	_ "github.com/kennek4/godo/cmd/group"
	_ "github.com/kennek4/godo/cmd/init"
	_ "github.com/kennek4/godo/cmd/list"
	_ "github.com/kennek4/godo/cmd/new"
	_ "github.com/kennek4/godo/cmd/task"
)

func main() {
	cmd.Execute()
}
