package main

import (
	"fmt"
	"os"
)

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println()
	fmt.Println("Usage: ")
	fmt.Println()
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Println(cmd.name + ": " + cmd.desc)
	}
	return nil
}
func commandMap(conf *config) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area/?limit20&offset="
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
	locations, err := handleMap(currUrl)
	if err != nil {
		return err
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}
