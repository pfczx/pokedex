package iternal

import (
	"fmt"
	"os"
	"github.com/pfczx/pokedex/api"
)

func CommandExit(*Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func CommandHelp(*Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage: ")
	fmt.Println()
	commands := GetCommands()
	for _, cmd := range commands {
		fmt.Println(cmd.name + ": " + cmd.desc)
	}
	return nil
}
func CommandMap(conf *Config) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area/?limit=20&offset="
	currUrl := ""
	if conf.nextUrl == "" {
		currUrl = baseUrl + fmt.Sprintf("%d", 0)
		conf.prevUrl = currUrl
		conf.nextUrl = baseUrl + fmt.Sprintf("%d", 20)
		conf.offset = 20
	} else {
		conf.offset += 20
		currUrl = baseUrl + fmt.Sprintf("%d", conf.offset)
		conf.prevUrl = currUrl
		conf.nextUrl = baseUrl + fmt.Sprintf("%d", conf.offset+20)

	}
	var locations []Location
	locations, err := HandleMap(currUrl)
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}
func commandMapb(conf *Config) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area/?limit=20&offset="
	currUrl := ""
	if conf.offset == 0 {
		fmt.Println("you are on the first page")
	} else {
		conf.offset -= 20
		currUrl = baseUrl + fmt.Sprintf("%d", conf.offset)
		conf.prevUrl = currUrl

	}
	var locations []Location
	locations, err := HandleMap(currUrl)
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}
