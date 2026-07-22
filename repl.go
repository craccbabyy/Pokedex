package main

import (
	"PokedexCLI/internal/pokeapi"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name   string
	desc   string
	callbk func(*config) error
}

type config struct {
	pokeapiClient pokeapi.Client
	NextURL       *string // url for next 20 locations
	PreviousURL   *string // url for last 20 locations
}

func startREPL(cfg *config) {
	//cfg := config{NextURL: baseURL}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ") // print prompt

		success := scanner.Scan() // waiting for user input
		if !success {
			break
		}
		rawInput := scanner.Text()      // grab text typed by user
		cleaned := cleanInput(rawInput) // clean using helper function from repl.go

		if len(cleaned) == 0 {
			continue
		}
		cmdName := cleaned[0] // check for "exit commands" using first word

		// now we have the command, we need to:
		// check if its in the MAP of SUPPORTED COMMANDS!
		cmd, exists := supportedCommands[cmdName]
		if !exists {
			fmt.Printf("Unknown command: %v\n", rawInput)
			continue
		}
		err := cmd.callbk(cfg)
		if err != nil {
			fmt.Printf("Error executing %s: %v\n", cmd.name, err)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading standard input:", err)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
