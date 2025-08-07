package main

import (
	"time"

	"github.com/marcossnikel/pokecli/cli"
	"github.com/marcossnikel/pokecli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(time.Second*5, time.Minute*5)
	cfg := cli.Config{
		PokeapiClient: pokeClient,
	}

	cli.StartReplCLI(&cfg)
}
