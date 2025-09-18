package iternal

type CliCommand struct {
	Name     string
	Desc     string
	Callback func(*Config, ...string) error
}
type Config struct {
	PrevUrl string
	NextUrl string
	Offset  int
}

func GetCommands() map[string]CliCommand {
	Commands := map[string]CliCommand{
		"exit": {
			Name:     "exit",
			Desc:     "exit pokedex",
			Callback: CommandExit,
		},
		"help": {
			Name:     "help",
			Desc:     "print desc of cmd",
			Callback: CommandHelp,
		},
		"map": {
			Name:     "map",
			Desc:     "display 0-20 locations, next call dispplay 20-40 and so on",
			Callback: CommandMap,
		},
		"mapb": {
			Name:     "mapb",
			Desc:     "similar to map but backwards",
			Callback: CommandMapb,
		},
		"explore": {
			Name:     "explore",
			Desc:     "display pokemons in location",
			Callback: CommandExplore,
		},
		"catch": {
			Name:     "catch",
			Desc:     "catch pokemon with chance of succes based on pokemon lvl",
			Callback: CommandCatch,
		},
		"inspect": {
			Name:     "inspect",
			Desc:     "inspect cathed pokemon",
			Callback: CommandInspect,
		},
		"pokedex": {
			Name:     "pokedex",
			Desc:     "list all of ur pokemon names",
			Callback: CommandPokedex,
		},
	}
	return Commands
}
