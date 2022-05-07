package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/business"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/controller"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/repository"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/service"
	"github.com/spf13/viper"
)

func PokemonRoutes(r *gin.Engine) *gin.Engine {
	pokemonFile := viper.Get("CSVFile").(string)

	pokemon := controller.NewPokemonController(
		business.NewPokemonBusiness(
			repository.NewPokemonRepository(pokemonFile),
			service.NewExternalPokemonAPI(),
		),
	)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/pokemons", pokemon.GetAllPokemons)
		v1.GET("/pokemons/:pokemonId", pokemon.GetPokemonByID)
		v1.POST("/pokemons/:pokemonId", pokemon.StorePokemonByID)
	}
	return r
}
