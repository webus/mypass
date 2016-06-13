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
	Short: "Add new password or login to MyPass",
	Long: "Add new password or login to MyPass",
	Run: func(cmd *cobra.Command, args []string) {
		if addLogin {
			if len(args) == 1 {
				fmt.Println("Pls, define value for login")
			} else if len(args) == 2{
				db.UpdateDataBucket("login", args[0], args[1])
				fmt.Println("Got it!")
			} else {
				fmt.Println("Pls, define key and login for key")
			}
		}else if len(args) == 0 {
			fmt.Println("Pls, define key for login")
			fmt.Println("eg: mypass get amazon")
			fmt.Println("and you get your login for Amazon")
		} else if len(args) == 1 {
			pass := base.GetEditorText()
			db.UpdateDataBucket("pass",args[0], pass)
			fmt.Println("Got it!")
		} else {
			fmt.Println("Pls, define only one key for login")
		}
	},
}

