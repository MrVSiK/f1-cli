package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var standingCmd = &cobra.Command{
	Use:   "std",
	Short: "Get driver/constructor standings",
	Long:  "This command fetches driver/constructor standings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Standing")
	},
}

func init() {
	rootCmd.AddCommand(standingCmd)
}
