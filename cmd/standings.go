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

var constructor string

// var l *widgets.List = widgets.NewList()

var standingCmd = &cobra.Command{
	Use:   "std",
	Short: "Get driver/constructor standings",
	Long:  "This command fetches driver/constructor standings",
	Run: func(cmd *cobra.Command, args []string) {
		if year < "1950" {
			fmt.Println("Invalid year")
			return
		}
		if err := ui.Init(); err != nil {
			log.Fatalf("failed to initialize termui: %v", err)
		}
		defer ui.Close()
		if constructor != "" {
			table := set_constructors(get_constructor_standings(args))
			footer := widgets.NewParagraph()
			footer.Text = "Press q to quit"
			footer.Title = "Keys"
			footer.SetRect(5, 5, 40, 15)
			footer.BorderStyle.Fg = ui.ColorYellow
			grid1 := ui.NewGrid()
			termwidth1, termheight1 := ui.TerminalDimensions()
			grid1.SetRect(0, 0, termwidth1, termheight1)
			grid1.Set(ui.NewRow(0.75, ui.NewCol(1.0, table)), ui.NewRow(0.25, ui.NewCol(1.0, footer)))
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
		} else {
			table1, table2 := set_drivers(get_drivers_standings(args))
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
		}
	},
}

func init() {
	standingCmd.Flags().StringVarP(&round, "round", "r", "", "Finishing order based on the given round number")
	standingCmd.Flags().StringVarP(&year, "year", "y", "", "Year of the given race")
	standingCmd.Flags().StringVarP(&constructor, "constructor", "c", "", "Get constructors standings")
	rootCmd.AddCommand(standingCmd)
}

func get_drivers_standings(args []string) util.Response {
	data_in_response := util.Response{}
	if round == "" {
		_, err := os.Stat(fmt.Sprintf("cache/drivers/%s_full.json", year))
		if err == nil {
			file, err := os.Open(fmt.Sprintf("cache/drivers/%s_full.json", year))
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			byteValue, _ := ioutil.ReadAll(file)
			json.Unmarshal(byteValue, &data_in_response)
			return data_in_response
		}
		resp, err := http.Get(fmt.Sprintf("https://ergast.com/api/f1/%s/driverStandings.json", year))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(body, &data_in_response)
		go cache_driver_data(data_in_response)
		return data_in_response
	}
	_, err := os.Stat(fmt.Sprintf("cache/drivers/%s_%s.json", year, round))
	if err == nil {
		file, err := os.Open(fmt.Sprintf("cache/drivers/%s_%s.json", year, round))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		byteValue, _ := ioutil.ReadAll(file)
		json.Unmarshal(byteValue, &data_in_response)
		return data_in_response
	}
	resp, err := http.Get(fmt.Sprintf("https://ergast.com/api/f1/%s/%s/driverStandings.json", year, round))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &data_in_response)
	go cache_driver_data(data_in_response)
	return data_in_response
}

func cache_driver_data(data util.Response) {
	if round == "" {
		file, _ := json.MarshalIndent(data, "", " ")

		if err := ioutil.WriteFile(fmt.Sprintf("cache/drivers/%s_full.json", year), file, 0644); err != nil {
			log.Fatal(err)
		}
		return
	}
	file, _ := json.MarshalIndent(data, "", " ")

	if err := ioutil.WriteFile(fmt.Sprintf("cache/drivers/%s_%s.json", year, round), file, 0644); err != nil {
		log.Fatal(err)
	}
}

func get_constructor_standings(args []string) util.Response {
	data_in_response := util.Response{}
	if round == "" {
		_, err := os.Stat(fmt.Sprintf("cache/cons/%s_full.json", year))
		if err == nil {
			file, err := os.Open(fmt.Sprintf("cache/cons/%s_full.json", year))
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			byteValue, _ := ioutil.ReadAll(file)
			json.Unmarshal(byteValue, &data_in_response)
			return data_in_response
		}
		resp, err := http.Get(fmt.Sprintf("https://ergast.com/api/f1/%s/constructorStandings.json", year))
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		json.Unmarshal(body, &data_in_response)
		go cache_constructor_data(data_in_response)
		return data_in_response
	}
	_, err := os.Stat(fmt.Sprintf("cache/cons/%s_%s.json", year, round))
	if err == nil {
		file, err := os.Open(fmt.Sprintf("cache/cons/%s_%s.json", year, round))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		byteValue, _ := ioutil.ReadAll(file)
		json.Unmarshal(byteValue, &data_in_response)
		return data_in_response
	}
	resp, err := http.Get(fmt.Sprintf("https://ergast.com/api/f1/%s/%s/constructorStandings.json", year, round))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(body, &data_in_response)
	go cache_constructor_data(data_in_response)
	return data_in_response
}

func cache_constructor_data(data util.Response) {
	if round == "" {
		file, _ := json.MarshalIndent(data, "", " ")

		if err := ioutil.WriteFile(fmt.Sprintf("cache/cons/%s_full.json", year), file, 0644); err != nil {
			log.Fatal(err)
		}
		return
	}
	file, _ := json.MarshalIndent(data, "", " ")

	if err := ioutil.WriteFile(fmt.Sprintf("cache/cons/%s_%s.json", year, round), file, 0644); err != nil {
		log.Fatal(err)
	}
}

func set_drivers(response util.Response) (*widgets.Table, *widgets.Table) {
	table1 := widgets.NewTable()
	table2 := widgets.NewTable()
	length := len(response.MRData.StandingsTable.StandingsLists[0].DriverStandings)
	first := length / 2
	second := length - first
	arr1 := make([][]string, first+1)
	arr2 := make([][]string, second+1)
	arr1[0] = []string{"Position", "Name", "Constructor", "Points"}
	arr2[0] = []string{"Position", "Name", "Constructor", "Points"}
	for i := 0; i < first; i++ {
		name := fmt.Sprintf("%s %s", response.MRData.StandingsTable.StandingsLists[0].DriverStandings[i].Driver.GivenName, response.MRData.StandingsTable.StandingsLists[0].DriverStandings[i].Driver.FamilyName)
		arr1[i+1] = []string{response.MRData.StandingsTable.StandingsLists[0].DriverStandings[i].Position, name, response.MRData.StandingsTable.StandingsLists[0].DriverStandings[i].Constructors[0].Name, response.MRData.StandingsTable.StandingsLists[0].DriverStandings[i].Points}
	}
	for i := first; i < length; i++ {
		name := fmt.Sprintf("%s %s", response.MRData.StandingsTable.StandingsLists[0].DriverStandings[i].Driver.GivenName, response.MRData.StandingsTable.StandingsLists[0].DriverStandings[i].Driver.FamilyName)
		arr2[i+1-first] = []string{response.MRData.StandingsTable.StandingsLists[0].DriverStandings[i].Position, name, response.MRData.StandingsTable.StandingsLists[0].DriverStandings[i].Constructors[0].Name, response.MRData.StandingsTable.StandingsLists[0].DriverStandings[i].Points}
	}
	table1.Rows = arr1
	table2.Rows = arr2
	return table1, table2
}

func set_constructors(response util.Response) *widgets.Table {
	table := widgets.NewTable()
	length := len(response.MRData.StandingsTable.StandingsLists[0].ConstructorStandings)
	arr1 := make([][]string, length+1)
	arr1[0] = []string{"Position", "Name", "Points"}
	for i := 0; i < length; i++ {
		arr1[i+1] = []string{response.MRData.StandingsTable.StandingsLists[0].ConstructorStandings[i].Position, response.MRData.StandingsTable.StandingsLists[0].ConstructorStandings[i].Constructor.Name, response.MRData.StandingsTable.StandingsLists[0].ConstructorStandings[i].Points}
	}
	table.Rows = arr1
	return table
}
