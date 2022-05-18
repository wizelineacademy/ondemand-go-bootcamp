package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/business"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
)

type pokemon struct {
	pokemonBusiness business.PokemonBusiness
	baseController
}

func NewPokemonController(pokemonBusiness business.PokemonBusiness) *pokemon {
	return &pokemon{
		pokemonBusiness: pokemonBusiness,
	}
}

// Get all pokemons
func (ctrl pokemon) GetAllPokemons(c *gin.Context) {
	pokemons, err := ctrl.pokemonBusiness.GetAll()

	if err != nil {
		ctrl.ResponseError(c, err)
	}

	ctrl.ResponseSucess(c, http.StatusOK, pokemons)
}

// GetPokemonByID get pokemon based on ID
func (ctrl pokemon) GetPokemonByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("pokemonId"))
	if err != nil {
		ctrl.ResponseError(c, model.NewURLParameterDoesNotFound("pokemonId"))
		return
	}
	pokemon, errBusiness := ctrl.pokemonBusiness.GetByID(id)
	if errBusiness != nil {
		ctrl.ResponseError(c, errBusiness)
		return
	}

	ctrl.ResponseSucess(c, http.StatusOK, pokemon)
}

// StorePokemonByID get pokemon based on ID
func (ctrl pokemon) StorePokemonByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("pokemonId"))
	if err != nil {
		ctrl.ResponseError(c, model.NewURLParameterDoesNotFound("pokemonId"))
		return
	}
	pokemon, errBusiness := ctrl.pokemonBusiness.StoreByID(id)
	if errBusiness != nil {
		ctrl.ResponseError(c, errBusiness)
		return
	}

	ctrl.ResponseSucess(c, http.StatusCreated, pokemon)
}
