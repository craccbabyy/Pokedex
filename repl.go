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
	callbk func(*config, ...string) error // added for args to be passed
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
		cmdName := cleaned[0]

		// adding arguments
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		// check if its in the MAP of SUPPORTED COMMANDS!
		cmd, exists := supportedCommands[cmdName]
		if !exists {
			fmt.Printf("Unknown command: %v\n", rawInput)
			continue
		}
		err := cmd.callbk(cfg, args...) // added args to be passed
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
