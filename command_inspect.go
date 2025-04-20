package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, arguments ...string) error {
	if len(arguments) == 0 {
		return errors.New("no arguments provided")
	}

	pokemon, exists := cfg.caughtPokemon[arguments[0]]
	if !exists {
		errMsg := fmt.Sprintf("have not caught %s yet \n", arguments[0])
		return errors.New(errMsg)
	}

	fmt.Printf("Name: %s \n", pokemon.Name)
	fmt.Printf("Height: %v \n", pokemon.Height)
	fmt.Printf("Weight: %v \n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v \n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, val := range pokemon.Types {
		fmt.Printf("  -%s \n", val.Type.Name)
	}

	return nil
}
