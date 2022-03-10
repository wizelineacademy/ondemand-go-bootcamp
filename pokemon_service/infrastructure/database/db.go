package db

import (
	"encoding/csv"
	"log"
	"os"
)

type Store interface {
	ReadData() ([][]string, error)
}

func NewStore(path string) Store {
	return &store{path}
}

type store struct {
	path string
}

func (d store) ReadData() ([][]string, error) {
	file, err := os.Open(d.path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return records, nil
}
