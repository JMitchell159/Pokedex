package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"

	"github.com/JMitchell159/pokedex/internal/pokeapi"
)

func commandCatch(cfg *config, pokemon string) error {
	if cfg.pokedex == nil {
		cfg.pokedex = make(map[string]pokeapi.Pokemon)
	}

	_, ok := cfg.pokedex[pokemon]
	if ok {
		return fmt.Errorf("you have already caught %s", pokemon)
	}

	pokeapiClient := cfg.pokeapiClient
	resp := pokeapi.Pokemon{}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	dat, ok := cfg.cache.Get(pokemon)
	if !ok {
		temp, err := pokeapiClient.ListPokemonInfo(pokemon)
		if err != nil {
			return err
		}

		resp = temp

		entry, err := json.Marshal(resp)
		if err != nil {
			return err
		}
		cfg.cache.Add(pokemon, entry)
	} else {
		err := json.Unmarshal(dat, &resp)
		if err != nil {
			return err
		}
	}

	floatExp := float64(resp.BaseExperience)
	numToBeat := 0.01 * (40 + math.Sqrt(5 * floatExp + 5))
	if numToBeat < rand.Float64() {
		fmt.Printf("%s was caught!\n", pokemon)
		cfg.pokedex[pokemon] = resp
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}