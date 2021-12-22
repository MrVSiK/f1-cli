package util

type Details struct {
	Number   string `json:"number"`
	Position string `json:"position"`
	Points   string `json:"points"`
	Driver   struct {
		GivenName  string `json:"givenName"`
		FamilyName string `json:"familyName"`
	} `json:"Driver"`
	Constructor struct {
		Name string `json:"name"`
	} `json:"Constructor"`
	Constructors []struct {
		Name string `json:"name"`
	} `json:"Constructors"`
}

type Race struct {
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
	Date    string    `json:"date"`
	Time    string    `json:"time"`
	Results []Details `json:"Results"`
}

type Table struct {
	Season string `json:"season"`
	Races  []Race `json:"Races"`
}

type stdList struct {
	DriverStandings      []Details `json:"DriverStandings"`
	ConstructorStandings []struct {
		Position    string `json:"position"`
		Points      string `json:"points"`
		Constructor struct {
			Name string `json:"name"`
		} `json:"Constructor"`
	} `json:"ConstructorStandings"`
}

type Standing struct {
	Season         string    `json:"season"`
	StandingsLists []stdList `json:"StandingsLists"`
}

type Data struct {
	Xmlns          string   `json:"xmlns"`
	Series         string   `json:"series"`
	Url            string   `json:"url"`
	Limit          string   `json:"limit"`
	Offset         string   `json:"offset"`
	Total          string   `json:"total"`
	RaceTable      Table    `json:"RaceTable"`
	StandingsTable Standing `json:"StandingsTable"`
}

type Response struct {
	MRData Data `json:"MRData"`
}
