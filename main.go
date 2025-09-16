package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/pfczx/pokedex/iternal"
)

func repl() error {
	var conf iternal.Config
	commands := iternal.GetCommands()
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\033[32m"+"Pokedex > "+ "\033[0m")
		if sc.Scan() {
			text := strings.TrimSpace(sc.Text())
			if cmd, exists := commands[text]; exists {
				if err := cmd.Callback(&conf); err != nil {
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
