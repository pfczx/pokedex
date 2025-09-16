package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() error {
	commands := GetCommands()
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
