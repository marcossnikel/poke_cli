package cli

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp(cfg *Config) error {
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
func commandMap(cfg *Config) error {
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
func commandMapBack(cfg *Config) error {
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
