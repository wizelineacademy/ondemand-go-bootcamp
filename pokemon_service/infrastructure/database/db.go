package db

import (
	"encoding/csv"
	"log"
	"os"
)

// Store is an interface for Store
type Store interface {
	ReadData() ([][]string, error)
}

// NewStore returns a new Store
func NewStore(path string) Store {
	return &store{path}
}

// store is a struct that implements the Store interface
type store struct {
	path string
}

// ReadData reads the data from the csv file
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
