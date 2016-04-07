package main

import (
	"os"
	"fmt"
	"mypass/cmd"
	"mypass/db"
)

func main() {
	db.InitDatabase()
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
