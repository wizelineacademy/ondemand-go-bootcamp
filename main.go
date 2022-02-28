package main

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"

	"github.com/PasHdez/ondemand-go-bootcamp/config"
	"github.com/PasHdez/ondemand-go-bootcamp/domain/model"
	"github.com/PasHdez/ondemand-go-bootcamp/infrastructure/datastore"
	"github.com/PasHdez/ondemand-go-bootcamp/infrastructure/router"
	"github.com/PasHdez/ondemand-go-bootcamp/registry"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()

	db.AutoMigrate(&model.Pokemon{})
	//db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Database.Server.Address)
	if err := e.Start(":" + config.C.Database.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
