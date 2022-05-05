package main

import (
	"net/http"

	"github.com/GabrielRendonP/ondemand-go-bootcamp/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/pokemons", controllers.GetPokemons)
	http.HandleFunc("/pokemon", controllers.GetPokemon)

	_ = http.ListenAndServe(":8080", nil)
}
