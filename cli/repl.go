package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marcossnikel/pokecli/internal/pokeapi"
)

type Config struct {
	PokeapiClient   pokeapi.Client
	nextPageURL     *string
	previousPageURL *string
}

func StartReplCLI(cfg *Config) {
	userInputReader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		userInputReader.Scan()
		words := cleanInput(userInputReader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if err := command.callback(cfg, args...); err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type CLICommand struct {
	name        string
	description string
	callback    func(cfg *Config, args ...string) error
}

func getCommands() map[string]CLICommand {
	return map[string]CLICommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays 20 locations of pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations of pokemon world",
			callback:    commandMapBack,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"explore": {
			name:        "explore <location-name>",
			description: "Explore an location to see which pokemons lives there",
			callback:    commandExplore,
		},
	}
}
