package main

import (
	"fmt"
	"os"
)

func init() { // func auto runs at start
	supportedCommands["exit"] = cliCommand{
		name:   "exit",
		desc:   "Exit the program",
		callbk: commandExit,
	}
}

func commandExit(cfg *config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
