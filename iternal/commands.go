package iternal

import (
	"fmt"
	"github.com/pfczx/pokedex/api"
	"os"
)

func CommandExit(conf *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func CommandHelp(conf *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage: ")
	fmt.Println()
	commands := GetCommands()
	for _, cmd := range commands {
		fmt.Println(cmd.Name + ": " + cmd.Desc)
	}
	return nil
}
func CommandMap(conf *Config, args ...string) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area/?limit=20&offset="
	currUrl := ""
	if conf.NextUrl == "" {
		currUrl = baseUrl + fmt.Sprintf("%d", 0)
		conf.PrevUrl = currUrl
		conf.NextUrl = baseUrl + fmt.Sprintf("%d", 20)
		conf.Offset = 20
	} else {
		conf.Offset += 20
		currUrl = baseUrl + fmt.Sprintf("%d", conf.Offset)
		conf.PrevUrl = currUrl
		conf.NextUrl = baseUrl + fmt.Sprintf("%d", conf.Offset+20)

	}
	var locations []api.Location
	locations, err := api.HandleMap(currUrl)
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}
func CommandMapb(conf *Config, args ...string) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area/?limit=20&offset="
	currUrl := ""
	if conf.Offset == 0 {
		fmt.Println("you are on the first page")
		return nil
	} else {
		conf.Offset -= 20
		currUrl = baseUrl + fmt.Sprintf("%d", conf.Offset)
		conf.PrevUrl = currUrl

	}
	var locations []api.Location
	locations, err := api.HandleMap(currUrl)
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}

func CommandExplore(conf *Config, args ...string) error {
	if len(args)==0{
    return fmt.Errorf("no argument")
	}
	url := "https://pokeapi.co/api/v2/location-area/"+args[0]
	pokemons, err := api.HandleExplore(url)
	if err != nil {
		return  err
	}
	for _, pokemon := range pokemons {
		fmt.Println(" - " + pokemon.Name)
	}
	return nil
}
