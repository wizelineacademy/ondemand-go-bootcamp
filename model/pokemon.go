package model

import (
	"errors"
	"github.com/gocarina/gocsv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Pokemon struct {
	Id             int     `json:"id" csv:"ID"`
	Name           string  `json:"name" csv:"Name"`
	Form           string  `json:"form" csv:"Form"`
	Type1          string  `json:"firstType" csv:"Type1"`
	Type2          string  `json:"secondType" csv:"Type2"`
	Total          float64 `json:"total" csv:"Total"`
	Hp             float64 `json:"hearthPoints" csv:"HP"`
	Attack         float64 `json:"attackPoints" csv:"Attack"`
	Defense        float64 `json:"defensePoints" csv:"Defense"`
	SpecialAttack  float64 `json:"specialAttackPoints" csv:"Sp.Atk"`
	SpecialDefense float64 `json:"specialDefensePoints" csv:"Sp.Def"`
	Speed          float64 `json:"speedPoints" csv:"Speed"`
	Generation     float64 `json:"generation" csv:"Generation"`
}

func GetPokemonById(id int) (*Pokemon, error) {
	var pokemon *Pokemon
	pokemonList, err := getPokemonCsv()

	if err != nil {
		return nil, err
	}

	for i, _ := range pokemonList {
		if pokemonList[i].Id == id {
			pokemon = pokemonList[i]
		}
	}
	if pokemon == nil {
		log.Println("Unable to find pokemon with id ", id)
		err = errors.New("unable to find pokemon")
	}
	return pokemon, err
}

func getPokemonCsv() ([]*Pokemon, error) {
	in, err := os.Open(getFileName("TEST"))
	if err != nil {
		log.Println(err)
	}
	defer in.Close()

	var pokemonList []*Pokemon

	if err := gocsv.UnmarshalFile(in, &pokemonList); err != nil {
		log.Println(err)
		return nil, err
	}
	return pokemonList, err
}

func getFileName(env string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	switch env {
	case "TEST":
		return os.Getenv("TEST_FILE")
	case "LARGE":
		return os.Getenv("LARGE_FILE")
	default:
		return os.Getenv("TEST_FILE")
	}
}
