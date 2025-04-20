package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
	"github.com/AidanRJ1/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	needArguments bool
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	caughtPokemon map[string]pokeapi.Pokemon
	nextLocationsURL *string
	previousLocationsURL *string
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			needArguments: false,
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			needArguments: false,
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays a list of the next 20 locations",
			needArguments: false,
			callback:    commandMapf,
		},
		"mapb": {
			name: "mapb",
			description: "Displays a list of the previous 20 locations",
			needArguments: false,
			callback: commandMapb,
		},
		"explore": {
			name: "explore",
			description: "Displays a list of pokemon found in provided area",
			needArguments: true,
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Try to catch provided pokemon",
			needArguments: true,
			callback: commandCatch,
		},
	}
	return commands
}

func cleanInput(text string) []string {
	formattedString := strings.ToLower(text)
	words := strings.Fields(formattedString)
	return words
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		var arguments []string
		if len(words) > 1 {
			arguments = words[1:]
		}

		commandName := words[0]
	

		if command, exists := getCommands()[commandName]; exists {
			if command.needArguments {
				err := command.callback(cfg, arguments...)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				err := command.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
			}
			continue
		} else {
			fmt.Println("Unknown Command")
			continue
		}
	}
}