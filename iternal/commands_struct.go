package iternal

type CliCommand struct {
	name     string
	desc     string
	callback func(*Config) error
}
type Config struct {
	prevUrl string
	nextUrl string
	offset int
}

func GetCommands() map[string]CliCommand {
	commands := map[string]CliCommand{
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
		"mapb":{
			name : "mapb",
			desc: "similar to map but backwards",
			callback: commandMapb,
		},
	}
	return commands
}
