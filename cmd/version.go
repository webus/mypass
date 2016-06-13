package cmd

import "fmt"
import "github.com/spf13/cobra"

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number of MyPass",
	Long: "All software has versions. This is MyPass's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Okay!", args)
	},
}
