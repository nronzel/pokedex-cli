package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, params []string) error {
	fmt.Println("Exiting the Pokedex..")
	os.Exit(0)
	return nil
}
