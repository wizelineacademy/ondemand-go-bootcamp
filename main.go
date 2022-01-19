package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	c "wizelineluisfarfan/ondemand-go-bootcamp/controller"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/go/pokemon/all", c.GetAllPokemon).Methods(http.MethodGet)
	router.HandleFunc("/go/pokemon", c.GetSpecificPokemon).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}
