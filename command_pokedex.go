package main

import (
	"fmt"
)

func init() {
	supportedCommands["pokedex"] = cliCommand{
		name:   "pokedex",
		desc:   "List all Pokemon you have caught",
		callbk: commandPokedex,
	}
}

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You have not caught any Pokemon yet")
	}
	fmt.Println("Your Pokedex:")
	for _, pokeMon := range cfg.caughtPokemon {
		fmt.Printf("  - %v\n", pokeMon.Name)
	}
	return nil
}
