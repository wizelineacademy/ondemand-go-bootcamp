package business

import (
	"log"

	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/repository"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/service"
)

//go:generate mockery --name PokemonBusiness --filename pokemon.go --outpkg mocks --structname PokemonBusinessMock --disable-version-string
type PokemonBusiness interface {
	GetAll() ([]*model.PokemonDTO, *model.ErrorHandler)
	GetByID(id int) (*model.PokemonDTO, *model.ErrorHandler)
	StoreByID(id int) (*model.PokemonDTO, *model.ErrorHandler)
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
func (s pokemonBusiness) GetAll() ([]*model.PokemonDTO, *model.ErrorHandler) {
	log.Println("enter to get all pokemons!!!")
	pokemons, err := s.pokemonRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

// GetByID get pokemon by his id
func (s pokemonBusiness) GetByID(id int) (*model.PokemonDTO, *model.ErrorHandler) {
	log.Println("enter to get pokemon by id!!!")
	pokemon, err := s.pokemonRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return pokemon, nil
}

// StoreByID get pokemon by his id
func (s pokemonBusiness) StoreByID(id int) (*model.PokemonDTO, *model.ErrorHandler) {
	log.Println("enter to search and store pokemon by id!!!")
	pokemonAPI, err := s.serviceAPI.GetPokemonFromAPI(id)
	if err != nil {
		return nil, err
	}
	pokemon, errRepository := s.pokemonRepository.StoreToCSV(*pokemonAPI)
	if err != nil {
		return nil, errRepository
	}
	return pokemon, nil
}
