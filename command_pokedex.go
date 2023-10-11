package main

import "fmt"

func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:")
	for pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon)
	}
	return nil
}
