package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GabrielRendonP/ondemand-go-bootcamp/helpers"
)

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	fmt.Println(query, id)
	pokemon, err := helpers.FindPokemon(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message":"Invalid pokemon id"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(pokemon)
	w.Write(response)
}
