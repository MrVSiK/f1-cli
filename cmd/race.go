package cmd

import (
	"encoding/json"
	util "f1-cli/util"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/spf13/cobra"
)

var round string

// var name string

var raceCmd = &cobra.Command{
	Use:   "race",
	Short: "Get race data",
	Long:  "This command fetches the data for a given race",
	Run: func(cmd *cobra.Command, args []string) {
		if round < "0" {
			fmt.Println("Invalid Round number")
			return
		}
		if year < "1950" {
			fmt.Println("Invalid year")
			return
		}
		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()
		data := get_race(args)
		fmt.Println("Get Data")
		table1, table2 := set_race(data)
		fmt.Println("Set Data")
		footer := widgets.NewParagraph()
		footer.Text = "Press q to quit"
		footer.Title = "Keys"
		footer.SetRect(5, 5, 40, 15)
		footer.BorderStyle.Fg = ui.ColorYellow
		grid1 := ui.NewGrid()
		termwidth1, termheight1 := ui.TerminalDimensions()
		grid1.SetRect(0, 0, termwidth1, termheight1)
		grid1.Set(ui.NewRow(0.75, ui.NewCol(0.5, table1), ui.NewCol(0.5, table2)), ui.NewRow(0.25, ui.NewCol(1.0, footer)))
		ui.Render(grid1)
		uiEvents := ui.PollEvents()
		for {
			select {
			case e := <-uiEvents:
				switch e.ID {
				case "q", "<C-c>":
					return
				case "<Resize>":
					payload := e.Payload.(ui.Resize)
					grid1.SetRect(0, 0, payload.Width, payload.Height)
					ui.Clear()
					ui.Render(grid1)
				}
			}
		}
	},
}

func init() {
	raceCmd.Flags().StringVarP(&round, "round", "r", "", "Finishing order based on the given round number")
	raceCmd.Flags().StringVarP(&year, "year", "y", "", "Year of the given race")
	rootCmd.AddCommand(raceCmd)
}

func get_race(args []string) util.Response {
	data_in_response := util.Response{}
	_, err := os.Stat(fmt.Sprintf("cache/race/%s_%s.json", year, round))
	if err == nil {
		file, err := os.Open(fmt.Sprintf("cache/race/%s_%s.json", year, round))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		byteValue, _ := ioutil.ReadAll(file)
		json.Unmarshal(byteValue, &data_in_response)
		return data_in_response
	}
	resp, err := http.Get(fmt.Sprintf("https://ergast.com/api/f1/%s/%s/results.json", year, round))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &data_in_response)
	go cache_race_data(data_in_response)
	return data_in_response
}

func set_race(response util.Response) (*widgets.Table, *widgets.Table) {
	table1 := widgets.NewTable()
	table2 := widgets.NewTable()
	length := len(response.MRData.RaceTable.Races[0].Results)
	first := length / 2
	second := length - first
	arr1 := make([][]string, first)
	arr2 := make([][]string, second)
	arr1[0] = []string{"Position", "Name", "Constructor", "Points"}
	arr2[0] = []string{"Position", "Name", "Constructor", "Points"}
	for i := 0; i < length/2; i++ {
		name := fmt.Sprintf("%s %s", response.MRData.RaceTable.Races[0].Results[i].Driver.GivenName, response.MRData.RaceTable.Races[0].Results[i].Driver.FamilyName)
		arr1[i] = []string{response.MRData.RaceTable.Races[0].Results[i].Position, name, response.MRData.RaceTable.Races[0].Results[i].Constructor.Name, response.MRData.RaceTable.Races[0].Results[i].Points}
	}
	for i := length/2 + 1; i < length; i++ {
		name := fmt.Sprintf("%s %s", response.MRData.RaceTable.Races[0].Results[i].Driver.GivenName, response.MRData.RaceTable.Races[0].Results[i].Driver.FamilyName)
		arr2[i-(length/2+1)] = []string{response.MRData.RaceTable.Races[0].Results[i].Position, name, response.MRData.RaceTable.Races[0].Results[i].Constructor.Name, response.MRData.RaceTable.Races[0].Results[i].Points}
	}
	table1.Rows = arr1
	table2.Rows = arr2
	return table1, table2
}

func cache_race_data(data util.Response) {
	file, _ := json.MarshalIndent(data, "", " ")

	if err := ioutil.WriteFile(fmt.Sprintf("cache/race/%s_%s.json", year, round), file, 0644); err != nil {
		log.Fatal(err)
	}
}
