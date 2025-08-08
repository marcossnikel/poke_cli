package cli

import (
	"errors"
	"fmt"
	"math/rand"
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

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...", name)
	fmt.Println()

	pokemon, err := cfg.PokeapiClient.FetchPokemon(name)
	if pokemon.Name == "" {
		fmt.Println("Invalid name provided, please type a valid pokemon name")
		return nil
	}
	if err != nil {
		return err
	}
	catch := rand.Intn(pokemon.BaseExperience) > 40
	msg := " escaped!"
	if !catch {
		fmt.Println(pokemon.Name + msg)
		return nil
	}
	cfg.CaughtPokemon[pokemon.Name] = pokemon
	msg = " was caught!"
	fmt.Println(pokemon.Name + msg)
	fmt.Println("You may now inspect with the inspect command.")

	return nil
}

func commandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...", name)
	fmt.Println()

	pokemon, hasCaught := cfg.CaughtPokemon[name]
	if !hasCaught {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}
	return nil
}

func commandPokedex(cfg *Config, args ...string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.CaughtPokemon) == 0 {
		fmt.Println("You have 0 pokeons caught!")
		return nil
	}
	for _, value := range cfg.CaughtPokemon {
		fmt.Printf("- %s\n", value.Name)
	}
	return nil
}
