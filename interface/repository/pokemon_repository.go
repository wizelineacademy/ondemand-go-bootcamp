package repository

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"

	"github.com/alexis-wizeline/ondemand-go-bootcamp/domain/model"
	repo "github.com/alexis-wizeline/ondemand-go-bootcamp/usecase/repository"
)

type pokemonRepository struct {
}

func NewPokemonRepository() repo.PokemonRepository {
	return pokemonRepository{}
}

func (p pokemonRepository) GetPokemons() ([]*model.Pokemon, error) {
	pokemons, err := getAllPokemons()

	if err != nil {
		return pokemons, err
	}

	return pokemons, nil
}

func (p pokemonRepository) GetPokemonById(id uint64) (*model.Pokemon, error) {
	pokemons, err := getAllPokemons()

	if err != nil {
		return nil, err
	}

	for _, pokemon := range pokemons {
		if id == pokemon.ID {
			return pokemon, nil
		}
	}

	return nil, nil
}

func getAllPokemons() ([]*model.Pokemon, error) {
	var pokemons []*model.Pokemon
	rows, err := openAndGetCSVData()

	if err != nil {
		return pokemons, err
	}

	pokemons, err = transformRowsToPokemons(pokemons, rows)

	if err != nil {
		return pokemons, err
	}

	return pokemons, nil

}

func transformRowsToPokemons(pokemons []*model.Pokemon, rows [][]string) ([]*model.Pokemon, error) {
	for _, row := range rows {
		pokemon := new(model.Pokemon)
		stringId := row[0]
		id, err := strconv.ParseUint(stringId, 10, 32)

		if err != nil {
			return nil, errors.New("invalid csv format")
		}

		pokemon.ID = id
		pokemon.Name = row[1]
		pokemon.Type = row[2]

		pokemons = append(pokemons, pokemon)

	}

	return pokemons, nil
}

func openAndGetCSVData() ([][]string, error) {
	f, err := os.OpenFile("./data/Pokemon.csv", os.O_RDONLY, 0755)

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(f)
	reader.Read()

	rows, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return rows, nil
}
