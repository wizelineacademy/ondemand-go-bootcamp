package repository

import (
	"log"
	"strconv"

	db "pas.com/v1/infrastructure"
	model "pas.com/v1/models"
)

// type pokemonRepository struct {
// 	Pokemons []model.Pokemon
// }

// type PokemonRepository interface {
// 	ReadDb(p *pokemonRepository) []model.Pokemon
// }

// func NewPokemonRepository(db *[]model.Pokemon) PokemonRepository {
// 	return &pokemonRepository{db}
// }

func convertToStruct(records [][]string) []model.Pokemon {
	var result []model.Pokemon
	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}
		var tmp model.Pokemon = model.Pokemon{Id: id, Name: record[1]}

		result = append(result, tmp)
	}
	return result
}

// func (p *pokemonRepository) ReadDb() {
// 	records := db.ReadData()
// 	p.Pokemons = convertToStruct(records)
// }

func GetPokemons() []model.Pokemon {
	records := db.ReadData()
	return convertToStruct(records)
}

func GetPokemon(id int) model.Pokemon {
	data := db.ReadData()
	pokemons := convertToStruct(data)
	for _, pokemon := range pokemons {
		if pokemon.Id == id {
			return pokemon
		}
	}
	return model.Pokemon{}
}

// func (pr *pokemonRepository) FindAll(p []model.Pokemon) []model.Pokemon {
// 	//var p pokemons
// 	pr.ReadDb()
// 	return pr.Pokemons
// }
