package mapper

import (
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
)

func PokemonAPItoPokemonCSV(pokemonAPI model.PokemonAPI) model.PokemonCSV {
	var type1 string
	var type2 string
	if len(pokemonAPI.Types) < 2 {
		type1 = pokemonAPI.Types[0].Type.Name
		type2 = ""
	} else {
		type1 = pokemonAPI.Types[0].Type.Name
		type2 = pokemonAPI.Types[1].Type.Name
	}
	return model.PokemonCSV{
		Id:             pokemonAPI.Id,
		Name:           pokemonAPI.Name,
		Height:         pokemonAPI.Height,
		Weight:         pokemonAPI.Weight,
		BaseExperience: pokemonAPI.BaseExperience,
		PrimaryType:    type1,
		SecondaryType:  type2,
	}
}
func PokemonDTOToPokemonCSV(pokemonDTO *model.PokemonDTO) model.PokemonCSV {
	return model.PokemonCSV{
		Id:             pokemonDTO.Id,
		Name:           pokemonDTO.Name,
		Height:         pokemonDTO.Height,
		Weight:         pokemonDTO.Weight,
		BaseExperience: pokemonDTO.BaseExperience,
		PrimaryType:    pokemonDTO.PrimaryType,
		SecondaryType:  pokemonDTO.SecondaryType,
	}
}
func PokemonAPIToPokemonDTO(pokemonAPI model.PokemonAPI) *model.PokemonDTO {
	var type1 string
	var type2 string
	if len(pokemonAPI.Types) < 2 {
		type1 = pokemonAPI.Types[0].Type.Name
		type2 = ""
	} else {
		type1 = pokemonAPI.Types[0].Type.Name
		type2 = pokemonAPI.Types[1].Type.Name
	}
	return &model.PokemonDTO{
		Id:             pokemonAPI.Id,
		Name:           pokemonAPI.Name,
		Height:         pokemonAPI.Height,
		Weight:         pokemonAPI.Weight,
		BaseExperience: pokemonAPI.BaseExperience,
		PrimaryType:    type1,
		SecondaryType:  type2,
	}
}

func PokemonCSVToPokemonDTO(pokemonCSV model.PokemonCSV) *model.PokemonDTO {
	return &model.PokemonDTO{
		Id:             pokemonCSV.Id,
		Name:           pokemonCSV.Name,
		Height:         pokemonCSV.Height,
		Weight:         pokemonCSV.Weight,
		BaseExperience: pokemonCSV.BaseExperience,
		PrimaryType:    pokemonCSV.PrimaryType,
		SecondaryType:  pokemonCSV.SecondaryType,
	}
}

func PokemonsCSVToPokemonsDTO(pokemonsCSV []model.PokemonCSV) []*model.PokemonDTO {
	pokemonsDTO := make([]*model.PokemonDTO, 0)
	for _, pokemonCSV := range pokemonsCSV {
		pokemonsDTO = append(pokemonsDTO, PokemonCSVToPokemonDTO(pokemonCSV))
	}
	return pokemonsDTO
}
