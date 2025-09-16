package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"pokedex/iternal/commands"
	"pokedex/iternal/commands_struct"
)

func repl() error {
	var config config
	conf := &config	
	commands := getCommands()
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		if sc.Scan() {
			text := strings.TrimSpace(sc.Text())
			if cmd, exists := commands[text]; exists {
				if err := cmd.callback(conf); err != nil {
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
