/*
Copyright © 2024 Ken Tabanay kentabanay@gmail.com
*/
package main

import (
	"log"

	"github.com/kennek4/genv"
	"github.com/kennek4/godo/cmd"
	_ "github.com/kennek4/godo/cmd/delete"
	_ "github.com/kennek4/godo/cmd/group"
	_ "github.com/kennek4/godo/cmd/init"
	_ "github.com/kennek4/godo/cmd/list"
	_ "github.com/kennek4/godo/cmd/new"
	_ "github.com/kennek4/godo/cmd/task"
	"github.com/kennek4/godo/internal/util/gddb"
)

func main() {
	cmd.Execute()

	err := genv.Save()
	if err != nil {
		log.Fatal(err)
	}

	gddb.Close(gddb.GodoDb)
}
