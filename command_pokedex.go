package main

import "fmt"

func commandPokedex(cfg *config, s string) error {
	fmt.Println("Your Pokedex:")

	if cfg.pokedex == nil {
		fmt.Println(" - Oh No! It looks like you have not caught any Pokemon")
		return nil
	}

	for poke := range cfg.pokedex {
		fmt.Printf(" - %s\n", poke)
	}

	return nil
}