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
	"github.com/spf13/cobra"
)

var constructor string

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
		fmt.Println("Standing")
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
