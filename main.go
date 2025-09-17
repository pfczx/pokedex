package main

import (
	"bufio"
	"fmt"
	"github.com/pfczx/pokedex/iternal"
	"os"
	"strings"
)

func repl() error {
	var conf iternal.Config
	commands := iternal.GetCommands()
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\033[32m" + "Pokedex > " + "\033[0m")
		if sc.Scan() {
			text := strings.TrimSpace(sc.Text())
			parts := strings.Fields(text)
			if len(parts) == 0 {
				continue
			}
			commandName := parts[0]
			var args []string
			if len(parts) > 1 {
				args = parts[1:]
			}
			if cmd, exists := commands[commandName]; exists {
				if err := cmd.Callback(&conf, args...); err != nil {
					fmt.Println(err)
					continue
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
