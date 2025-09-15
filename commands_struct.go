package main

type cliCommand struct {
	name     string
	desc     string
	callback func() error
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
	}
	return commands
}
