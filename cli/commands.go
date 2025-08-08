package cli

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(cfg *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
func commandMap(cfg *Config, args ...string) error {
	locations, err := cfg.PokeapiClient.ListLocations(cfg.nextPageURL)

	if err != nil {
		return err
	}
	cfg.nextPageURL = locations.Next
	cfg.previousPageURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
func commandMapBack(cfg *Config, args ...string) error {
	if cfg.previousPageURL == nil {
		return errors.New("you're on the first page")
	}
	locations, err := cfg.PokeapiClient.ListLocations(cfg.previousPageURL)
	if err != nil {
		return err
	}

	cfg.nextPageURL = locations.Next
	cfg.previousPageURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	name := args[0]
	location, err := cfg.PokeapiClient.ListLocationByName(name)
	if err != nil {
		return err
	}

	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *Config, args ...string)
