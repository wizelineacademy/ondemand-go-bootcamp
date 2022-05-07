package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/business"
)

type pokemon struct {
	pokemonBusiness business.PokemonBusiness
}

type Pokemon interface {
	GetAllPokemons()
}

func NewPokemonController(pokemonBusiness business.PokemonBusiness) *pokemon {
	return &pokemon{
		pokemonBusiness: pokemonBusiness,
	}
}
func (p pokemon) GetAllPokemons(c *gin.Context) {
	pokemons, err := p.pokemonBusiness.GetAll()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, pokemons)
}

// GetPokemonByID get pokemon based on ID
func (p pokemon) GetPokemonByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("pokemonId"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	pokemon, err := p.pokemonBusiness.GetByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, pokemon)
}

// StorePokemonByID get pokemon based on ID
func (p pokemon) StorePokemonByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("pokemonId"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	pokemon, err := p.pokemonBusiness.StoreByID(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, pokemon)
}
