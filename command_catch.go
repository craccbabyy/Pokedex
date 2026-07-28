package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

func init() {
	supportedCommands["catch"] = cliCommand{
		name:   "catch <pokemon_name>",
		desc:   "Attempt to catch a Pokemon by throwing a Pokeball at it",
		callbk: commandCatch,
	}
}

// LIST THE POKEMON in the given LOCATION
func commandCatch(cfg *config, args ...string) error {
	// https://pokeapi.co/api/v2/pokedex/{id or name}/

	// first validate that only one pokemon arg was supplied
	if len(args) != 1 {
		return errors.New("please provide a single Pokemon name")
	}
	pokemonName := args[0]

	// call the new method on the client
	catchResp, err := cfg.pokeapiClient.CatchPokemon(pokemonName) //////////////////
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", catchResp.Name)

	// now we need to use "BaseExperience" along with 'math/rand' to make a chance factor
	// higher BaseExperience value means higher difficulty
	luck := rand.IntN(250) + 1
	if luck >= catchResp.BaseExperience { // if the random number is high enough, catch the pokemon
		fmt.Printf("%v was caught!\n", catchResp.Name)
	} else {
		fmt.Printf("%v escaped!\n", catchResp.Name)
	}
	cfg.caughtPokemon[catchResp.Name] = catchResp
	return nil
}

/*
output should look like:
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu escaped!
Pokedex > catch pikachu
Throwing a Pokeball at pikachu...
pikachu was caught!
*/
