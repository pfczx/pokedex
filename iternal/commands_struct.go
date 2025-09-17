package iternal

type CliCommand struct {
	Name     string
	Desc     string
	Callback func(*Config,...string) error
}
type Config struct {
	PrevUrl string
	NextUrl string
	Offset int
}

func GetCommands() map[string]CliCommand {
	Commands := map[string]CliCommand{
		"exit": {
			Name:     "exit",
			Desc:     "exit pokedex",
			Callback: CommandExit,
		},
		"help":{
			Name: "help",
			Desc: "print desc of cmd",
			Callback: CommandHelp,
		},
		"map":{
			Name: "map",
			Desc: "display 0-20 locations, next call dispplay 20-40 and so on",
			Callback: CommandMap,
		},
		"mapb":{
			Name : "mapb",
			Desc: "similar to map but backwards",
			Callback: CommandMapb,
		},
		"explore":{
			Name: "explore",
			Desc: "display pokemons in location",
			Callback: CommandExplore,
		},
	}
	return Commands
}
