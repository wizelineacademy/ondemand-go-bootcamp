package repository

import (
	"os"

	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model/mapper"

	"github.com/gocarina/gocsv"
)

//go:generate mockery --name PokemonRepository --filename pokemon.go --outpkg mocks --structname PokemonRepositoryMock --disable-version-string
type PokemonRepository interface {
	GetAll() ([]*model.PokemonDTO, *model.ErrorHandler)
	GetByID(id int) (*model.PokemonDTO, *model.ErrorHandler)
	StoreToCSV(pokemonAPI model.PokemonAPI) (*model.PokemonDTO, *model.ErrorHandler)
	GetCSVDataInMemory() (map[int]model.PokemonCSV, *model.ErrorHandler)
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
func (p *pokemonRepository) GetAll() ([]*model.PokemonDTO, *model.ErrorHandler) {
	pokemonFile, err := p.openFile()
	if err != nil {
		return nil, err
	}
	pokemons := []model.PokemonCSV{}

	if err := gocsv.UnmarshalFile(pokemonFile, &pokemons); err != nil {
		return nil, model.NewUnmarshalFileError(err.Error())
	}
	defer p.closeFile(pokemonFile)
	return mapper.PokemonsCSVToPokemonsDTO(pokemons), nil
}

// openFile open the csv file
func (p pokemonRepository) openFile() (*os.File, *model.ErrorHandler) {
	filePokemon, err := os.OpenFile(p.file, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, model.NewOpenFileError(err.Error())
	}
	return filePokemon, nil
}

// closeFile close csv file
func (p pokemonRepository) closeFile(file *os.File) {
	file.Close()
}

// GetByID get pokemon from csv by id
func (p pokemonRepository) GetByID(id int) (*model.PokemonDTO, *model.ErrorHandler) {
	pokemons, err := p.GetAll()
	if err != nil {
		return nil, err
	}
	for _, pokemon := range pokemons {
		if pokemon.Id == id {
			return pokemon, nil
		}
	}

	return nil, model.NewNotFoundPokemonError(id)
}

// StoreToCSV store pokemon to csv
func (p pokemonRepository) StoreToCSV(pokemonAPI model.PokemonAPI) (*model.PokemonDTO, *model.ErrorHandler) {
	pokemonMap, err := p.GetCSVDataInMemory()
	if err != nil {
		return nil, err
	}
	pokemon := mapper.PokemonAPItoPokemonCSV(pokemonAPI)
	pokemonMap[pokemon.Id] = pokemon
	pokemons := make([]model.PokemonCSV, 0)
	for _, pokemonObj := range pokemonMap {
		pokemons = append(pokemons, pokemonObj)
	}
	pokemonFile, err := p.openFile()
	if err != nil {
		return nil, err
	}
	if err := gocsv.MarshalFile(&pokemons, pokemonFile); err != nil {
		return nil, model.NewAccesingCSVFileError(err.Error())
	}
	p.closeFile(pokemonFile)
	return mapper.PokemonAPIToPokemonDTO(pokemonAPI), nil
}

// getCSVDataInMemory store pokemons from csv to memory
func (p pokemonRepository) GetCSVDataInMemory() (map[int]model.PokemonCSV, *model.ErrorHandler) {
	pokemonMap := make(map[int]model.PokemonCSV)
	pokemons, err := p.GetAll()
	if err != nil {
		return nil, err
	}
	for _, pokemon := range pokemons {
		pokemonMap[pokemon.Id] = mapper.PokemonDTOToPokemonCSV(pokemon)
	}
	return pokemonMap, nil
}
