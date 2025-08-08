package main

import (
	"fmt"
	"os"
	"net/http"
	"encoding/json"
	"io"
)

type result struct {
	name	string	`json:"name"`
	url		string	`json:"url"`
}

type config struct {
	next		string		`json:"next"`
	previous	string		`json:"previous"`
	results		[]result	`json:"results"`
}

type cliCommand struct {
	name		string
	description	string
	callback	func() error
	configPtr	*config
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func help(cliMap map[string]cliCommand) func() error {
	return func() error {
		fmt.Print("Welcome to the Pokedex!\nUsage\n\n")
		for _, val := range cliMap {
			fmt.Printf("%s: %s\n", val.name, val.description)
		}
		return nil
	}
}

func commandMap(configPtr *config) func() error {
	return func() error {
		req, err := http.NewRequest("GET", configPtr.next, nil)
		if err != nil {
			return fmt.Errorf("Couldn't make the request struct: %v\n", err)
		}

		client := &http.Client{}

		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		res, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("Something wrong with the response: %v\n", err)
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)

		if res.StatusCode > 299 {
			return fmt.Errorf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
		}

		if err != nil {
			return fmt.Errorf("Error reading response body: %v\n", err)
		}

		fmt.Printf("Response Body (string): %s\n", string(body))

		err = json.Unmarshal(body, configPtr)
		if err != nil {
			return err
		}
		for _, result := range configPtr.results {
			fmt.Printf("%s\n", result.name)
		}
		return nil
	}
}

func mapb(configPtr *config) func() error {
	return func() error {
		if configPtr.previous == "null" {
			fmt.Print("you're on the first page\n")
		} else {
			res, err := http.Get(configPtr.previous)
			if err != nil {
				return err
			}
			body, err := io.ReadAll(res.Body)
			res.Body.Close()
			if res.StatusCode > 299 {
				return fmt.Errorf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
			}
			if err != nil {
				return err
			}
			dat := []byte(body)
			err = json.Unmarshal(dat, configPtr)
			if err != nil {
				return err
			}
			for _, result := range configPtr.results {
				fmt.Printf("%s\n", result.name)
			}
		}
		return nil
	}
}
