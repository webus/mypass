package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "mypass",
	Short: "MyPass is a very fast password manager",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
