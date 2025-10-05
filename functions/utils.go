package functions

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const url = "https://groupietrackers.herokuapp.com/api/artists"

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
	Locations []string `json:"locations"`
}

type Date struct {
	Dates []string `json:"dates"`
}

type Relation struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type All struct {
	Artist   Artist
	Location Location
	Dates    Date
	Relation Relation
}

type ErrorPage struct {
	Code    int
	Message string
}

func FetchJson(url string, target any) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to get url'sdata")
	}

	defer resp.Body.Close()

	err2 := json.NewDecoder(resp.Body).Decode(target)
	if err2 != nil {
		return fmt.Errorf("failed to fetch data")
	}

	return nil
}

func GetDetails(url string) (All, error) {
	var artist Artist
	err := FetchJson(url, &artist)
	if err != nil {
		return All{}, err
	}

	var location Location
	err1 := FetchJson(artist.Locations, &location)
	if err1 != nil {
		return All{}, err1
	}

	var dates Date
	err2 := FetchJson(artist.ConcertDates, &dates)
	if err2 != nil {
		return All{}, err2
	}

	var relation Relation
	err3 := FetchJson(artist.Relations, &relation)
	if err3 != nil {
		return All{}, err3
	}

	return All{
		Artist:   artist,
		Location: location,
		Dates:    dates,
		Relation: relation,
	}, nil
}
