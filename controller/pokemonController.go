package controller

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"wizelineluisfarfan/ondemand-go-bootcamp/model"
	"wizelineluisfarfan/ondemand-go-bootcamp/util"
)

func GetAllPokemon(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	log.Println("Retrieving all pokemon in the DB")
	// Code

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func GetSpecificPokemon(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	pokemonId, err := strconv.Atoi(req.URL.Query()["pokemonId"][0])
	if err != nil{
		util.ThrowBadRequestError(w, errors.New("request is expecting an int value"))
		return
	}
	log.Println("Retrieving pokemon", pokemonId)

	var response []byte
	pokemon, err := model.GetPokemonById(pokemonId)
	if err != nil{
		util.ThrowError(w, err)
		return
	}
	response, _ = json.Marshal(&pokemon)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}


