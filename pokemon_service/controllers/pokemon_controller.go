package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	repository "pas.com/v1/services"
)

type controller struct {
	repository repository.PokemonRepository
}

type Controller interface {
	GetPokemon(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetPokemons(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	GetPokemonsPars(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

func NewPokemonController(repository repository.PokemonRepository) Controller {
	return &controller{repository}
}

func (c *controller) GetPokemons(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	poks := c.repository.GetPokemons()
	json.NewEncoder(w).Encode(poks)
}

func (c *controller) GetPokemon(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pokemon := c.repository.GetPokemon(idInt)
	json.NewEncoder(w).Encode(pokemon)
}

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
