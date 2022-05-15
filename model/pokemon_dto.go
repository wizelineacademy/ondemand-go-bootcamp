package model

type PokemonDTO struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"baseExperience"`
	PrimaryType    string `json:"primaryType"`
	SecondaryType  string `json:"secondaryType"`
}
