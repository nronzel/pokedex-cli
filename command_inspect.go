package main

import "fmt"

func commandInspect(cfg *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("you must provide a pokemon name")
	}

	name := args[0]

	if pokemon, ok := cfg.caughtPokemon[name]; ok {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)

		// Print stats
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}

		// Print types
		fmt.Println("Types:")
		for _, t := range pokemon.Types {
			fmt.Printf(" -%s\n", t.Type.Name)
		}

		return nil
	}

	fmt.Printf("you have not caught %s\n", name)
	return nil
}
