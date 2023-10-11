package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreArea(location string) ([]string, error) {
	url := baseURL + "/location-area/" + location

	// Check if data is in cache
	if val, ok := c.cache.Get(url); ok {
		var response LocationAreaResponse
		err := json.Unmarshal(val, &response)
		if err != nil {
			return nil, err
		}

		pokemonNames := parsePokemonNames(response)

		return pokemonNames, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response LocationAreaResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, data)

	pokemonNames := parsePokemonNames(response)

	return pokemonNames, nil

}

func parsePokemonNames(response LocationAreaResponse) []string {
	var pokemonNames []string
	for _, encounter := range response.PokemonEncounters {
		pokemonNames = append(pokemonNames, encounter.Pokemon.Name)
	}
	return pokemonNames
}
