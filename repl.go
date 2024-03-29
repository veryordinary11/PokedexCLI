package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" >")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}

		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "show this help menu",
			callback:    callbackHelp,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "catch a pokemon",
			callback:    callbackCatch,
		},
		"map": {
			name:        "map",
			description: "show a map of the region",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "show the next page of the map of the region",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "explore a location area",
			callback:    callbackExplore,
		},
		"exit": {
			name:        "exit",
			description: "exit the program",
			callback:    callbackExit,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "inspect a pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "show your pokedex",
			callback:    callbackPokedex,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
