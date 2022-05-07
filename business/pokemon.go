package business

import (
	"log"

	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/repository"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/service"
)

type PokemonBusiness interface {
	GetAll() ([]model.Pokemon, error)
	GetByID(id int) (*model.Pokemon, error)
	StoreByID(id int) (*model.Pokemon, error)
}

// PokemonService dependencies from Pokemon service
type pokemonBusiness struct {
	pokemonRepository repository.PokemonRepository
	serviceAPI        service.ExternalPokemonAPI
}

// NewPokemonService initializer method for create PokemonService
func NewPokemonBusiness(repository repository.PokemonRepository, service service.ExternalPokemonAPI) *pokemonBusiness {
	return &pokemonBusiness{
		pokemonRepository: repository,
		serviceAPI:        service,
	}
}

// GetAll get all pokemons from repository
func (s pokemonBusiness) GetAll() ([]model.Pokemon, error) {
	log.Println("Enter to get all pokemons!!!")
	pokemons, err := s.pokemonRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

// GetByID get pokemon by his id
func (s pokemonBusiness) GetByID(id int) (*model.Pokemon, error) {
	log.Println("Enter to get pokemon by id!!!")
	pokemon, err := s.pokemonRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// StoreByID get pokemon by his id
func (s pokemonBusiness) StoreByID(id int) (*model.Pokemon, error) {
	log.Println("Enter to search and store pokemon by id!!!")
	pokemonAPI, err := s.serviceAPI.GetPokemonFromAPI(id)
	if err != nil {
		return nil, err
	}
	pokemon, err := s.pokemonRepository.StoreToCSV(*pokemonAPI)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}
