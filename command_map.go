package main

import (
	"fmt"

	"github.com/nronzel/pokedex-cli/internal"
)

type Config struct {
	Previous string
	Next     string
}

func commandMap() error {
	data, err := api.FetchLocation()
	if err != nil {
		fmt.Println(err)
	}
	results := data.Results
	for _, result := range results {
		fmt.Println(result.Name)
	}
	return nil
}
