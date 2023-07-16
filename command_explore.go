package main

import (
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	pokeapiClient := cfg.pokeapiClient

	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, got %v", len(args))
	}

	locationName := args[0]

	res, err := pokeapiClient.GetLocationArea(locationName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:\n", locationName)
	for _, pokemon := range res.PokemonEncounters {
		fmt.Println("- ", pokemon.Pokemon.Name)
	}

	return nil
}
