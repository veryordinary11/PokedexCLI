package main

import (
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("expected 1 argument, got %v", len(args))
	}

	pokemonName := args[0]

	pokemon, ok := cfg.pokemons[pokemonName]
	if !ok {
		return fmt.Errorf("you don't have %s", pokemonName)
	}

	fmt.Printf("Inspecting %s:\n", pokemonName)
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Base Experience: %v\n", pokemon.BaseExperience)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Abilities: \n")
	for _, ability := range pokemon.Abilities {
		fmt.Printf("\t%s\n", ability.Ability.Name)
	}
	fmt.Printf("Moves: \n")
	for _, move := range pokemon.Moves {
		fmt.Printf("\t%s\n", move.Move.Name)
	}

	return nil
}
