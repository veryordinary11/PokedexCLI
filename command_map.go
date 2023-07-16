package main

import (
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	pokeapiClient := cfg.pokeapiClient

	res, err := pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range res.Results {
		fmt.Println("- ", area.Name)
	}

	cfg.nextLocationAreaURL = res.Next
	cfg.previousLocationAreaURL = res.Previous

	return nil
}

func callbackMapb(cfg *config, args ...string) error {
	pokeapiClient := cfg.pokeapiClient

	if cfg.previousLocationAreaURL == nil {
		return fmt.Errorf("no previous location area")
	}

	res, err := pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range res.Results {
		fmt.Println("- ", area.Name)
	}

	cfg.nextLocationAreaURL = res.Next
	cfg.previousLocationAreaURL = res.Previous

	return nil
}
