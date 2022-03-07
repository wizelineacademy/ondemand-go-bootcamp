package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	repository "pas.com/v1/services"
)

func GetPokemons(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	//var p model.Pokemon
	pokemons := repository.GetPokemons()
	json.NewEncoder(w).Encode(pokemons)
}

func GetPokemon(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	id := ps.ByName("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pokemon := repository.GetPokemon(idInt)
	json.NewEncoder(w).Encode(pokemon)
}
