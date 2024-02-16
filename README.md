# Pokédex cli

This is a small project in Go from  [boot.dev](https://www.boot.dev/) where you make a REPL Pokedex.

## To Install
```bash
go install 
```
- command `pokedexcli`

## To build

```bash
git clone 

cd pokedexcli

go build
./pokedexcli
```


### Commands implemented:
 - help: Lists all the available commands
 - catch {pokemon_name}: Attempt to catch a pokemon and add it to your pokedex
 - inspect {pokemon_name}: View info of a caught pokemon
 - pokedex: View all the pokemon in your pokedex
 - help: prints the help menu
 - exit: exits from pokodex shell
 - map: list the next page of location areas
 - mapb: list the previous page of location areas
 - explore {location_area}: lists the pokemon in location area
