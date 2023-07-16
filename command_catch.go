package main

import (
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	pokeapiClient := cfg.pokeapiClient

	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, got %v", len(args))
	}

	pokemonName := args[0]

	pokemon, err := pokeapiClient.CatchPokemon(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Try to catch %s:\n", pokemonName)

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(randNum, threshold, pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}

	_, ok := cfg.pokemons[pokemonName]
	if ok {
		return fmt.Errorf("you already have %s", pokemonName)
	}

	cfg.pokemons[pokemonName] = pokemon

	fmt.Printf("You caught %s!\n", pokemonName)

	return nil
}
