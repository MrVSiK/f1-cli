package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of f1 tool",
	Long:  "This command can be used get the version number of f1 tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deployer v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
