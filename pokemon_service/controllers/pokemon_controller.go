package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	repository "github.com/PasHdez/ondemand-go-bootcamp/repository"
	"github.com/julienschmidt/httprouter"
)

// controller is the implementation of Controller interface
type controller struct {
	repository repository.PokemonRepository
}

// Controller is the interface that provides the controller methods
type Controller interface {
	GetPokemon(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetPokemons(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetPokemonsPars(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

// NewController returns a new instance of Controller
func NewPokemonController(repository repository.PokemonRepository) Controller {
	return &controller{repository}
}

// GetPokemons returns a json response with all pokemons
func (c *controller) GetPokemons(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	poks, err := c.repository.GetPokemons()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(poks)
}

// GetPokemon returns a json response with a pokemon by id
func (c *controller) GetPokemon(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pokemon, err := c.repository.GetPokemon(idInt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(pokemon)
}

// GetPokemonsPars returns a json response with a list of pokemons by type and quantity
func (c *controller) GetPokemonsPars(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	pType := ps.ByName("pType")
	items, err := strconv.Atoi(ps.ByName("items"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	itemsPerWorker, err := strconv.Atoi(ps.ByName("itemsPerWorker"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pokemon, err := c.repository.GetPokemonsPars(pType, items, itemsPerWorker)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(pokemon)
}
