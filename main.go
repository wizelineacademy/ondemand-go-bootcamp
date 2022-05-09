package main

import (
	"flag"
	"github.com/alexis-wizeline/ondemand-go-bootcamp/infraestructure/router"
	"log"

	"github.com/labstack/echo/v4"
)

var addr = flag.String("addr", ":8000", "http service address")

func main() {
	flag.Parse()

	e := router.NewRouter(echo.New())

	log.Fatal(e.Start(*addr))
}
