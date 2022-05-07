package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/router/v1"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r = router.PokemonRoutes(r)

	r.Run()
}
