package repo

import (
	"encoding/csv"
	"log"
	"os"
)

type localData struct{}

func NewLocalData() localData {
	return localData{}
}

func (r localData) ReadCSVData() ([][]string, error) {
	var data [][]string

	file, err := os.Open("./lib/pokemon.csv")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err = csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return data, nil
}
