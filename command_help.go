package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Available commands:")

	availableCommands := getCommands()
	for _, command := range availableCommands {
		fmt.Println(command.name, "-", command.description)
	}

	fmt.Println("")

	return nil
}
