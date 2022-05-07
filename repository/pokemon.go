package repository

import (
	"errors"
	"os"

	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model/mapper"

	"github.com/gocarina/gocsv"
)

type PokemonRepository interface {
	GetAll() ([]model.Pokemon, error)
	GetByID(id int) (*model.Pokemon, error)
	StoreToCSV(pokemonAPI model.PokemonAPI) (*model.Pokemon, error)
	GetCSVDataInMemory() (map[int]model.Pokemon, error)
}

// PokemonRepository structure for repository, contains the csv file's name
type pokemonRepository struct {
	file string
}

// NewPokemonRepository method for create a Repository instance
func NewPokemonRepository(csvFilename string) *pokemonRepository {
	return &pokemonRepository{
		file: csvFilename,
	}
}

// GetAll get all pokemons from csv file
func (p *pokemonRepository) GetAll() ([]model.Pokemon, error) {
	pokemonFile, err := p.openFile()
	if err != nil {
		return nil, err
	}
	pokemons := []model.Pokemon{}

	if err := gocsv.UnmarshalFile(pokemonFile, &pokemons); err != nil {
		return nil, errors.New("There was a problem parsing the csv file")
	}
	defer p.closeFile(pokemonFile)
	return pokemons, nil
}

// openFile open the csv file
func (p pokemonRepository) openFile() (*os.File, error) {
	filePokemon, err := os.OpenFile(p.file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, errors.New("There was a problem opening the csv file")
	}
	return filePokemon, nil
}

// closeFile close csv file
func (p pokemonRepository) closeFile(file *os.File) {
	file.Close()
}

// GetByID get pokemon from csv by id
func (p pokemonRepository) GetByID(id int) (*model.Pokemon, error) {
	pokemons, err := p.GetAll()
	if err != nil {
		return nil, err
	}
	for _, pokemon := range pokemons {
		if pokemon.Id == id {
			return &pokemon, nil
		}
	}
	return nil, errors.New("the pokemon does not exist")
}

// StoreToCSV store pokemon to csv
func (p pokemonRepository) StoreToCSV(pokemonAPI model.PokemonAPI) (*model.Pokemon, error) {
	pokemonMap, err := p.GetCSVDataInMemory()
	if err != nil {
		return nil, err
	}
	pokemon := mapper.PokemonAPItoPokemon(pokemonAPI)
	pokemonMap[pokemon.Id] = pokemon
	pokemons := make([]model.Pokemon, 0)
	for _, pokemonObj := range pokemonMap {
		pokemons = append(pokemons, pokemonObj)
	}
	pokemonFile, err := p.openFile()
	if err != nil {
		return nil, err
	}
	if err := gocsv.MarshalFile(&pokemons, pokemonFile); err != nil {
		return nil, errors.New("There was a problem accesing to csv file")
	}
	p.closeFile(pokemonFile)
	return &pokemon, nil
}

// getCSVDataInMemory store pokemons from csv to memory
func (p pokemonRepository) GetCSVDataInMemory() (map[int]model.Pokemon, error) {
	pokemonMap := make(map[int]model.Pokemon)
	pokemons, err := p.GetAll()
	if err != nil {
		return nil, err
	}
	for _, pokemon := range pokemons {
		pokemonMap[pokemon.Id] = pokemon
	}
	return pokemonMap, nil
}
