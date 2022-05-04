package helpers

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"strconv"
)

type PokemonEntry struct {
	Number   int64
	Name     string
	PokeType string
}

func FetchCSV(url string) ([][]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	log.Println(response.Body)

	csvReader := csv.NewReader(response.Body)
	data, err := csvReader.ReadAll()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func ReadPokemonData() []PokemonEntry {
	var data [][]string

	file, err := os.Open("./lib/pokemon.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err = csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	var pokeList []PokemonEntry

	for i, line := range data {
		if i > 0 {
			var pokemon PokemonEntry
			for j, pokeAtt := range line {
				if j == 0 {
					pokemon.Number, _ = strconv.ParseInt(pokeAtt, 10, 0)
				} else if j == 1 {
					pokemon.Name = pokeAtt
				} else if j == 2 {
					pokemon.PokeType = pokeAtt
				}
			}
			pokeList = append(pokeList, pokemon)
		}
	}
	return pokeList
}
