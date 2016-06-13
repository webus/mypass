package cmd

import (
	"log"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/webus/mypass/db"
	"github.com/webus/mypass/base"
)

var getLogin bool

func init() {
	RootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().BoolVarP(&getLogin, "login", "l", false, "Use login")
}

var getCmd = &cobra.Command{
	Use: "get",
	Short: "Get password or login from MyPass",
	Long: "Get password or login from MyPass",
	Run: func(cmd *cobra.Command, args []string) {
		if getLogin {
			if len(args) == 0 {
				fmt.Println("Pls, define key")
			} else if len(args) == 1{
				data, err := db.GetDataBucket("login",args[0])
				if err != nil {
					log.Fatal(err)
				}
				base.WriteAll(data)
				fmt.Println("Check your clipboard.")
			}
		} else if len(args) == 0 {
			fmt.Println("Pls, define key for password")
			fmt.Println("eg: mypass get amazon")
			fmt.Println("and you get your password for Amazon")
		} else if len(args) == 1 {
			data, err := db.GetDataBucket("pass", args[0])
			if err != nil {
				log.Fatal(err)
			}
			// debug
			/*
			log.Println(args[0])
			log.Println("HERE")
			log.Println(string(data))
                        */
			// debug
			base.WriteAll(data)
			fmt.Println("Check your clipboard.")
		} else {
			fmt.Println("Pls, define only one key for password")
		}
	},
}
