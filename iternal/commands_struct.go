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
			callback: CommandExit,
		},
		"help":{
			name: "help",
			desc: "print desc of cmd",
			callback: CommandHelp,
		},
		"map":{
			name: "map",
			desc: "display 0-20 locations, next call dispplay 20-40 and so on",
			callback: CommandMap,
		},
		"mapb":{
			name : "mapb",
			desc: "similar to map but backwards",
			callback: CommandMapb,
		},
	}
	return commands
}
