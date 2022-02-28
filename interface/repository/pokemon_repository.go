package repository

import (
	"github.com/PasHdez/ondemand-go-bootcamp/domain/model"
	"github.com/jinzhu/gorm"
)

type pokemonRepository struct {
	db *gorm.DB
}

type PokemonRepository interface {
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
}

func NewPokemonRepository(db *gorm.DB) PokemonRepository {
	return &pokemonRepository{db}
}

func (pr *pokemonRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	err := pr.db.Find(&p).Error
	if err != nil {
		return nil, err
	}

	return p, nil
}
