/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package main

import (
	"github.com/kennek4/godo/cmd"
	_ "github.com/kennek4/godo/cmd/init"
	_ "github.com/kennek4/godo/cmd/new"
)

func main() {
	cmd.Execute()
}
