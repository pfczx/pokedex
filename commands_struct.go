package main

type cliCommand struct {
	name     string
	desc     string
	callback func(*config) error
}
type config struct {
	prevUrl string
	nextUrl string
	offset int
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:     "exit",
			desc:     "exit pokedex",
			callback: commandExit,
		},
		"help":{
			name: "help",
			desc: "print desc of cmd",
			callback: commandHelp,
		},
		"map":{
			name: "map",
			desc: "display 0-20 locations, next call dispplay 20-40 and so on",
			callback: commandMap,
		},
	}
	return commands
}
