package main

import (
	"encoding/json"
	"fmt"

	"github.com/JMitchell159/pokedex/internal/pokeapi"
)

func explore(cfg *config, location string) error {
	pokeapiClient := cfg.pokeapiClient

	resp := pokeapi.LocationAreaResp{}

	dat, ok := cfg.cache.Get(location)
	if !ok {
		temp, err := pokeapiClient.ListLocationInfo(location)
		if err != nil {
			return err
		}

		resp = temp

		entry, err := json.Marshal(resp)
		if err != nil {
			return err
		}
		cfg.cache.Add(location, entry)
	} else {
		err := json.Unmarshal(dat, &resp)
		if err != nil {
			return err
		}
	}

	fmt.Printf("Exploring %s...\n", location)
	fmt.Println("Found Pokemon:")
	for _, result := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", result.Pokemon.Name)
	}

	return nil
}
