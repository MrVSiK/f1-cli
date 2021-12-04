package cmd

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

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

var raceCmd = &cobra.Command{
	Use:   "race",
	Short: "Get race statistics",
	Long:  "This command fetches data about the given race",
	Run: func(cmd *cobra.Command, args []string) {
		data := getData()
	},
}

func init() {
	rootCmd.AddCommand(raceCmd)
}

func getData() response {
	data_in_response := response{}
	resp, err := http.Get("https://ergast.com/api/f1/2020.json")
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
