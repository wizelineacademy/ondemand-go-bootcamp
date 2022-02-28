package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/viper"
)

type config struct {
	Database struct {
		User                 string
		Password             string
		Net                  string
		Addr                 string
		DBName               string
		AllowNativePasswords bool
		Params               struct {
			ParseTime string
		}
		Server struct {
			Address string
		}
	}
}

var C config

func ReadConfig() {
	Config := &C
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	path := filepath.Join(os.Getenv("GOPATH"), "config/")
	viper.AddConfigPath(path) //filepath.Join("$GOPATH", "src", "github.com", "PasHdez", "ondemand-go-bootcamp", "config"))
	// fmt.Printf("trying to connect to: %v \n", path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		log.Fatalln(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	spew.Dump(C)
}
