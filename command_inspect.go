package main

import (
	"errors"
	"fmt"
)

func init() {
	supportedCommands["inspect"] = cliCommand{
		name:   "inspect <pokemon_name>",
		desc:   "Inspect a Pokemon you have caught",
		callbk: commandInspect,
	}
}

func commandInspect(cfg *config, args ...string) error {
	// this command is different, we are only accessing the cache (not making any http request)
	// and therefore has no associated method in pokeapi

	//check for exactly 1 argument
	if len(args) != 1 {
		return errors.New("please provide a Pokemon name")
	}
	name := args[0]
	//check the cache for the pokemon name in users caught pokemon
	pokemon, ok := cfg.caughtPokemon[name]
	if !ok {
		return errors.New("you have not caught that pokemon yet!")
	}

	// if we made it here, the pokemon was already caught, so return it's details
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  -%v\n", typeInfo.Type.Name)
	}
	return nil
}

/* desired output structure :
Pokedex > inspect pidgey
Name: pidgey
Height: 3
Weight: 18
Stats:
  -hp: 40
  -attack: 45
  -defense: 40
  -special-attack: 35
  -special-defense: 35
  -speed: 56
Types:
  - normal
  - flying
*/
