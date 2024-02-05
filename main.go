package main

import (
	"time"

	"github.com/GnarlyLasagna/go-repl-pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationAreasApi *string
	prevLocationAreasApi *string
	caughtPokemon        map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Minute * 5),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)

}
