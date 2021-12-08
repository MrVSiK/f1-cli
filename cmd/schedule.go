package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/cobra"
)

type race struct {
	Season   string `json:"season"`
	Round    string `json:"round"`
	Url      string `json:"url"`
	RaceName string `json:"raceName"`
	Circuit  struct {
		CircuitId   string `json:"circuitId"`
		Url         string `json:"url"`
		CircuitName string `json:"circuitName"`
		Location    struct {
			Lat      string `json:"lat"`
			Long     string `json:"long"`
			Locality string `json:"locality"`
			Country  string `json:"country"`
		} `json:"Location"`
	} `json:"Circuit"`
	Date string `json:"date"`
	Time string `json:"time"`
}

type table struct {
	Season string `json:"season"`
	Races  []race `json:"Races"`
}

type data struct {
	Xmlns     string `json:"xmlns"`
	Series    string `json:"series"`
	Url       string `json:"url"`
	Limit     string `json:"limit"`
	Offset    string `json:"offset"`
	Total     string `json:"total"`
	RaceTable table  `json:"RaceTable"`
}

type response struct {
	MRData data `json:"MRData"`
}

type table_data struct {
	name        string
	circuitName string
	round       string
	location    string
	date        string
	time        string
}

var t *widgets.Table = widgets.NewTable()

var year string

var scheduleCmd = &cobra.Command{
	Use:   "sch",
	Short: "Get race schedule for a season",
	Long:  "This command fetches the schedule for a given year",
	Run: func(cmd *cobra.Command, args []string) {
		if year < "1950" {
			fmt.Println("Invalid year")
			return
		}
		data := getData(args)
		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()
		footer := widgets.NewParagraph()
		footer.Text = "Press q to quit\nPress j to scroll down\nPress k to scroll up"
		footer.Title = "Keys"
		footer.SetRect(5, 5, 40, 15)
		footer.BorderStyle.Fg = ui.ColorYellow
		l := widgets.NewList()
		l.Title = "Schedule"
		l.Rows = writeData(data.MRData.RaceTable.Races)
		l.TextStyle = ui.NewStyle(ui.ColorYellow)
		l.WrapText = false
		l.SetRect(0, 0, 25, 8)
		grid := ui.NewGrid()
		termwidth, termheight := ui.TerminalDimensions()
		grid.SetRect(0, 0, termwidth, termheight)
		grid.Set(ui.NewRow(0.75, ui.NewCol(0.4, l), ui.NewCol(0.6, createDetails(getDetails(data.MRData.RaceTable.Races, l.SelectedRow)))), ui.NewRow(0.25, ui.NewCol(1.0, footer)))
		t.Title = "Details"
		ui.Render(grid)
		uiEvents := ui.PollEvents()
		for {
			select {
			case e := <-uiEvents:
				switch e.ID {
				case "q", "<C-c>":
					return
				case "<Resize>":
					payload := e.Payload.(ui.Resize)
					grid.SetRect(0, 0, payload.Width, payload.Height)
					ui.Clear()
					ui.Render(grid)
				case "j", "<Down>":
					l.ScrollDown()
					grid.Set(ui.NewRow(0.75, ui.NewCol(0.4, l), ui.NewCol(0.6, createDetails(getDetails(data.MRData.RaceTable.Races, l.SelectedRow)))), ui.NewRow(0.25, ui.NewCol(1.0, footer)))
					ui.Clear()
					ui.Render(grid)
				case "k", "<Up>":
					l.ScrollUp()
					grid.Set(ui.NewRow(0.75, ui.NewCol(0.4, l), ui.NewCol(0.6, createDetails(getDetails(data.MRData.RaceTable.Races, l.SelectedRow)))), ui.NewRow(0.25, ui.NewCol(1.0, footer)))
					ui.Clear()
					ui.Render(grid)
				}
			}
		}
	},
}

func init() {
	scheduleCmd.Flags().StringVarP(&year, "year", "y", "", "Schedule of the given year")
	rootCmd.AddCommand(scheduleCmd)
}

func getData(args []string) response {
	data_in_response := response{}
	resp, err := http.Get(fmt.Sprintf("https://ergast.com/api/f1/%s.json", year))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &data_in_response)
	return data_in_response
}

func writeData(data []race) []string {
	size := len(data)
	arr := make([]string, size)
	for i := 0; i < size; i++ {
		arr[i] = data[i].RaceName
	}
	return arr
}

func getDetails(data []race, current int) *table_data {
	return &table_data{
		name:        data[current].RaceName,
		circuitName: data[current].Circuit.CircuitName,
		round:       data[current].Round,
		location:    data[current].Circuit.Location.Country,
		date:        data[current].Date,
		time:        data[current].Time,
	}
}

func createDetails(data *table_data) *widgets.Table {
	arr := make([][]string, 6)
	arr[0] = []string{"Race:", data.name}
	arr[1] = []string{"Round Number:", data.round}
	arr[2] = []string{"Circuit Name:", data.circuitName}
	arr[3] = []string{"Country:", data.location}
	arr[4] = []string{"Date:", data.date}
	arr[5] = []string{"Time:", data.time}
	t.Rows = arr
	return t
}
