package main

import "fmt"

func commandExplore(cfg *config, params []string) error {
	if len(params) == 0 {
		return fmt.Errorf("Please provide an area name to explore. E.g., explore <area_name>")
	}
	areaName := params[0]
	pokemonNames, err := cfg.pokeapiClient.ExploreArea(areaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", areaName)
	for _, name := range pokemonNames {
		fmt.Printf(" - %s\n", name)
	}
	return nil

}
