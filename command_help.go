package main

import (
	"fmt"
)

func init() { // func auto runs at start
	supportedCommands["help"] = cliCommand{
		name:   "help",
		desc:   "Display a help message",
		callbk: commandHelp,
	}
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range supportedCommands {
		fmt.Printf("  %-10s: %s\n", cmd.name, cmd.desc)
	}
	return nil
}
