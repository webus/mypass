package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/webus/mypass/base"
	"github.com/webus/mypass/db"
)

var addLogin bool

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().BoolVarP(&addLogin, "login", "l", false, "Use login")
}

var addCmd = &cobra.Command{
	Use: "add",
	Short: "add",
	Long: "add",
	Run: func(cmd *cobra.Command, args []string) {
		if addLogin {
			if len(args) == 1 {
				fmt.Println("Pls, define value for login")
			} else if len(args) == 2{
				db.UpdateDataBucket("login",args[0], args[1])
			} else {
				fmt.Println("Pls, define key and login for key")
			}
		} else if len(args) == 1 {
			pass := base.GetEditorText()
			db.UpdateDataBucket("pass",args[0], pass)
		}
	},
}

