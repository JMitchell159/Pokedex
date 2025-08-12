package main

import (
	"errors"
	"fmt"
	"encoding/json"

	"github.com/JMitchell159/pokedex/internal/pokeapi"
)

func commandMap(cfg *config) error {
	pokeapiClient := cfg.pokeapiClient

	resp := pokeapi.LocationAreasResp{}

	if cfg.nextLocationAreaURL != nil {
		dat, ok := cfg.cache.Get(*cfg.nextLocationAreaURL)
		if !ok {
			temp, err := pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
			if err != nil {
				return err
			}

			resp = temp

			entry, err := json.Marshal(resp)
			if err != nil {
				return err
			}
			cfg.cache.Add(*cfg.nextLocationAreaURL, entry)
		} else {
			err := json.Unmarshal(dat, &resp)
			if err != nil {
				return err
			}
		}
	} else {
		temp, err := pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
		if err != nil {
			return err
		}

		resp = temp

		entry, err := json.Marshal(resp)
		if err != nil {
			return err
		}
		cfg.cache.Add("https://pokeapi.co/api/v2/location-area", entry)
	}

	fmt.Println("Location areas:")
	for _, result := range resp.Results {
		fmt.Printf(" - %s\n", result.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func commandMapb(cfg *config) error {
	pokeapiClient := cfg.pokeapiClient

	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}

	resp := pokeapi.LocationAreasResp{}

	dat, ok := cfg.cache.Get(*cfg.prevLocationAreaURL)
	if !ok {
		temp, err := pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
		if err != nil {
			return err
		}

		resp = temp

		entry, err := json.Marshal(resp)
		if err != nil {
			return err
		}
		cfg.cache.Add(*cfg.prevLocationAreaURL, entry)
	} else {
		err := json.Unmarshal(dat, &resp)
		if err != nil {
			return err
		}
	}

	fmt.Println("Location areas:")
	for _, result := range resp.Results {
		fmt.Printf(" - %s\n", result.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}
