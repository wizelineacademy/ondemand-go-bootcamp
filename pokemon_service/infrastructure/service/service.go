package service

import (
	"log"

	db "github.com/PasHdez/ondemand-go-bootcamp/infrastructure/database"
)

type service struct {
	Store db.Store
}

type Service interface {
	GetData() [][]string
}

func NewService(store db.Store) Service {
	return &service{store}
}

func (s *service) GetData() [][]string {
	result, err := s.Store.ReadData()
	if err != nil {
		log.Fatal(err)
	}
	return result
}
