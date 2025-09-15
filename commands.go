package main

import(
	"fmt"
	"os"

)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage: ")
	fmt.Println()
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Println(cmd.name + ": " + cmd.desc)
	}
	return nil
}


