package service

import (
	"errors"
	"strconv"

	db "github.com/PasHdez/ondemand-go-bootcamp/infrastructure/database"
	models "github.com/PasHdez/ondemand-go-bootcamp/models"
)

var (
	ErrPokemonNotFound  = errors.New("error: Pokemon not found")
	ErrPokemonsNotFound = errors.New("error: Pokemons not found")
)

type service struct {
	Store    db.Store
	Pokemons []models.Pokemon
}

type Service interface {
	GetAll() ([]models.Pokemon, error)
	GetById(Id int) (models.Pokemon, error)
}

func NewService(store db.Store) (Service, error) {
	srv := &service{Store: store}
	if err := srv.loadPokemons(); err != nil {
		return nil, err
	}
	return srv, nil
}

func (a *service) loadPokemons() error {
	data, err := a.Store.ReadData()
	if err != nil {
		return err
	}
	for _, row := range data {
		id, err := strconv.Atoi(row[0])
		if err != nil {
			return err
		}
		pokemon := models.Pokemon{
			Id:     id,
			Name:   row[1],
			Url:    row[2],
			Sprite: row[3],
		}
		a.Pokemons = append(a.Pokemons, pokemon)
	}
	return nil
}

func (s *service) GetAll() ([]models.Pokemon, error) {
	if len(s.Pokemons) == 0 {
		return nil, ErrPokemonsNotFound
	}

	return s.Pokemons, nil
}

func (s *service) GetById(Id int) (models.Pokemon, error) {
	for _, pokemon := range s.Pokemons {
		if pokemon.Id == Id {
			return pokemon, nil
		}
	}
	return models.Pokemon{}, ErrPokemonNotFound
}
