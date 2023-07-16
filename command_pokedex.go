package main

import "fmt"

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
