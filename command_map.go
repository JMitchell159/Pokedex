package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	pokeapiClient := cfg.pokeapiClient

	resp, err := pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
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

	resp, err := pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, result := range resp.Results {
		fmt.Printf(" - %s\n", result.Name)
	}
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}
