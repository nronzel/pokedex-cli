package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationArea struct {
	Count    int      `json:"count"`
	Next     string   `json:"next"`
	Previous *string  `json:"previous"` // pointer to allow Null string value
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func FetchLocation() (LocationArea, error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/location-area")
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var location LocationArea
	err = json.Unmarshal(body, &location)
	if err != nil {
		return LocationArea{}, err
	}
	return location, nil
}
