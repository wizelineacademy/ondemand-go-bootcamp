package helpers

import (
	"fmt"
	"log"

	"github.com/GabrielRendonP/ondemand-go-bootcamp/repo"
)

type PokemonEntry struct {
	Number   string
	Name     string
	PokeType string
}

func GetAllPokemons() ([]PokemonEntry, error) {
	var lr = repo.NewLocalData()
	data, err := lr.ReadCSVData()

	if err != nil {
		log.Panic("Could not read csv data")
		return nil, err
	}

	var pokeList []PokemonEntry
	for i, line := range data {
		if i > 0 {
			var pokemon PokemonEntry
			for j, pokeAtt := range line {
				if j == 0 {
					pokemon.Number = pokeAtt
				} else if j == 1 {
					pokemon.Name = pokeAtt
				} else if j == 2 {
					pokemon.PokeType = pokeAtt
				}
			}
			pokeList = append(pokeList, pokemon)
		}
	}
	return pokeList, nil
}

func FindPokemon(id string) (PokemonEntry, error) {
	pokeList, err := GetAllPokemons()

	if err != nil {
		return PokemonEntry{}, fmt.Errorf("error in data list")
	}

	for _, poke := range pokeList {
		if poke.Number == id {
			log.Println("Found!", poke)
			return poke, nil
		}
	}

	return PokemonEntry{}, fmt.Errorf("pokemon with id %s not found", id)
}
