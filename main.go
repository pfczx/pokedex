package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

//

func repl() error {
	commands := getCommands()
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")

		if sc.Scan() {
			text := strings.TrimSpace(sc.Text())
			if cmd, exists := commands[text]; exists {
				if err := cmd.callback(); err != nil {
					return err
				}

			} else {
				fmt.Println("Unknown command")
			}
			if err := sc.Err(); err != nil {
				fmt.Println("Error reading input:", err)
			}
		}
	}
}
func main() {
	if err := repl(); err != nil {
		fmt.Print(err)
	}
}
