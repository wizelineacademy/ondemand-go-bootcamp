package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/alexis-wizeline/ondemand-go-bootcamp/usecase/repository"
)

type pokemonController struct {
	pokemonRepository repository.PokemonRepository
}

type PokemonController interface {
	GetPokemons(c echo.Context) error
	GetPokemonById(c echo.Context) error
}

func NewPokemonController(pr repository.PokemonRepository) PokemonController {
	return pokemonController{pokemonRepository: pr}
}

func (p pokemonController) GetPokemonById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	pokemon, err := p.pokemonRepository.GetPokemonById(id)

	if pokemon == nil && err == nil {
		return c.JSON(http.StatusNotFound, "pokemon not found")
	} else if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, pokemon)

}

func (p pokemonController) GetPokemons(c echo.Context) error {
	pokemons, err := p.pokemonRepository.GetPokemons()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, pokemons)
}
