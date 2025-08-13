package main

import "fmt"

func commandInspect(cfg *config, pokemon string) error {
	dat, ok := cfg.pokedex[pokemon]
	if !ok {
		return fmt.Errorf("you have not caught %s", pokemon)
	}

	fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nSpecies: %s\nStats:\n", dat.Name, dat.Height, dat.Weight, dat.Species.Name)
	for _, val := range dat.Stats {
		fmt.Printf(" - %s: %d\n", val.Stat.Name, val.BaseStat)
	}

	fmt.Println("Types:")
	for _, val := range dat.Types {
		fmt.Printf(" - %s\n", val.Type.Name)
	}

	fmt.Println("Abilities:")
	for _, val := range dat.Abilities {
		fmt.Printf(" - %s\n", val.Ability.Name)
	}

	return nil
}
