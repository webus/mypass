package main

import (
	"os"
	"fmt"
	"github.com/webus/mypass/cmd"
	"github.com/webus/mypass/db"
)

func main() {
	db.InitDatabase()
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
