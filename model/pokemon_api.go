package model

type PokemonAPI struct {
	Id             int        `json:"id"`
	BaseExperience int        `json:"baseExperience"`
	Height         int        `json:"height"`
	Weight         int        `json:"weight"`
	Name           string     `json:"name"`
	Types          []TypeSlot `json:"types"`
}

type TypeSlot struct {
	Type Type `json:"type"`
}

type Type struct {
	Name string `json:"name"`
}
