package service

import (
	"errors"
	"log"
	"math"
	"strconv"
	"sync"

	db "github.com/PasHdez/ondemand-go-bootcamp/infrastructure/database"
	models "github.com/PasHdez/ondemand-go-bootcamp/models"
)

var (
	// ErrPokemonsNotFound is returned when no pokemon is found
	ErrPokemonNotFound = errors.New("error: Pokemon not found")
	// ErrPokemonsNotFoound is returned when no pokemons are found
	ErrPokemonsNotFound = errors.New("error: Pokemons not found")
	// ErrPokemonInvalidParameterType is returned when the parameter type is invalid
	ErrPokemonInvalidParameterType = errors.New("error: Pokemon invalid parameter type")
	// ErrMaxItemsPerWorkersReached is returned when the max items per worker is reached
	ErrMaxItemsPerWorkersReached = errors.New("error: Max items per workers quantity reached")
)

// MaxItemsPerWorkerQuantity is the max quantity of items per worker
const MaxItemsPerWorkerQuantity int = 10

// service is the implementation of Service interface
type service struct {
	Store    db.Store
	Pokemons []models.Pokemon
}

// Service is the interface that provides the service methods
type Service interface {
	GetAll() ([]models.Pokemon, error)
	GetById(Id int) (models.Pokemon, error)
	GetByPars(pType string, items int, itemsPerWorker int) ([]models.Pokemon, error)
}

// NewService returns a new instance of Service
func NewService(store db.Store) (Service, error) {
	srv := &service{Store: store}
	if err := srv.loadPokemons(); err != nil {
		return nil, err
	}
	return srv, nil
}

// loadPokemons loads the pokemons from the csv file
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

// GetAll returns all pokemons
func (s *service) GetAll() ([]models.Pokemon, error) {
	if len(s.Pokemons) == 0 {
		return nil, ErrPokemonsNotFound
	}

	return s.Pokemons, nil
}

// GetById returns a pokemon by id
func (s *service) GetById(Id int) (models.Pokemon, error) {
	for _, pokemon := range s.Pokemons {
		if pokemon.Id == Id {
			return pokemon, nil
		}
	}
	return models.Pokemon{}, ErrPokemonNotFound
}

// findPokemonWorker is the worker function that finds the pokemons by type
func findPokemonWorker(pType string, itemsPerWorker int, jobs <-chan models.Pokemon, results chan<- models.Pokemon) {
	log.Printf("Worker started with %v items", itemsPerWorker)
	found := 0
	for p := range jobs {
		if pType == "odd" && p.Id%2 != 0 {
			found++
			results <- p
		} else if pType == "even" && p.Id%2 == 0 {
			found++
			results <- p
		}
		if found == itemsPerWorker {
			log.Printf("worker of %v items finished", itemsPerWorker)
			return
		}
	}
}

// GetByPars returns a slice of pokemons by type
func (s *service) GetByPars(pType string, items int, itemsPerWorker int) ([]models.Pokemon, error) {
	if pType != "odd" && pType != "even" {
		return nil, ErrPokemonInvalidParameterType
	}

	if itemsPerWorker > MaxItemsPerWorkerQuantity {
		return nil, ErrMaxItemsPerWorkersReached
	}

	wCount := int(math.Ceil(float64(items) / float64(itemsPerWorker)))

	jobs := make(chan models.Pokemon, len(s.Pokemons))
	results := make(chan models.Pokemon, (len(s.Pokemons)/2)+1)

	wg := sync.WaitGroup{}
	rItems := items
	for w := 1; w <= wCount; w++ {
		actualItemsPerWorker := itemsPerWorker
		if rItems < itemsPerWorker {
			actualItemsPerWorker = rItems
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			findPokemonWorker(pType, actualItemsPerWorker, jobs, results)
		}()
		rItems -= itemsPerWorker
	}

	for _, pokemon := range s.Pokemons {
		jobs <- pokemon
	}

	close(jobs)
	wg.Wait()
	close(results)

	pokeResults := []models.Pokemon{}
	for result := range results {
		pokeResults = append(pokeResults, result)
	}
	return pokeResults, nil
}
