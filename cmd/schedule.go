package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var scheduleCmd = &cobra.Command{
	Use:   "sch",
	Short: "Get race schedule for a season",
	Long:  "This command fetches the schedule for a given year",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Schedule")
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
}
