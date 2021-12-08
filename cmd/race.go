package cmd

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/cobra"
)

var scheduleCmd = &cobra.Command{
	Use:   "race",
	Short: "Get race data",
	Long:  "This command fetches the data for a given race",
	Run: func(cmd *cobra.Command, args []string) {
		l := widgets.NewList()
		l.Title = "List"
		l.Rows = []string{
			"[0] github.com/gizak/termui/v3",
			"[1] [你好，世界](fg:blue)",
			"[2] [こんにちは世界](fg:red)",
			"[3] [color](fg:white,bg:green) output",
			"[4] output.go",
			"[5] random_out.go",
			"[6] dashboard.go",
			"[7] foo",
			"[8] bar",
			"[9] baz",
		}
		l.TextStyle = ui.NewStyle(ui.ColorYellow)
		l.WrapText = false
		l.SetRect(0, 0, 25, 8)

		ui.Render(l)

		previousKey := ""
		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "q", "<C-c>":
				return
			case "j", "<Down>":
				l.ScrollDown()
			case "k", "<Up>":
				l.ScrollUp()
			case "<C-d>":
				l.ScrollHalfPageDown()
			case "<C-u>":
				l.ScrollHalfPageUp()
			case "<C-f>":
				l.ScrollPageDown()
			case "<C-b>":
				l.ScrollPageUp()
			case "g":
				if previousKey == "g" {
					l.ScrollTop()
				}
			case "<Home>":
				l.ScrollTop()
			case "G", "<End>":
				l.ScrollBottom()
			}

			if previousKey == "g" {
				previousKey = ""
			} else {
				previousKey = e.ID
			}

			ui.Render(l)
		}
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
}
