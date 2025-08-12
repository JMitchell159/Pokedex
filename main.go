package main

import (
	"time"

	"github.com/JMitchell159/pokedex/internal/pokeapi"
	"github.com/JMitchell159/pokedex/internal/pokecache"
)

type config struct {
	pokeapiClient       pokeapi.Client
	cache               *pokecache.Cache
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
		cache: pokecache.NewCache(time.Minute),
	}

	startRepl(&cfg)
}
