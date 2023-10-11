package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonByName(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return pokemon, nil
}
