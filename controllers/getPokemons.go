package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/GabrielRendonP/ondemand-go-bootcamp/helpers"
)

func GetPokemons(w http.ResponseWriter, r *http.Request) {

	pokeList, err := helpers.ReadPokemonData()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"No list found"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(pokeList)
	w.Write(response)
}
