package main

import (
	"time"

	"github.com/veryordinary11/PokedexCLI/internal/pokeapi"
)

type config struct {
	pokeapiClient           pokeapi.Client
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
	pokemons                map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		pokemons:      make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
