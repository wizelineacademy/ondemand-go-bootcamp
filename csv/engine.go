package csv

import (
	"encoding/csv"
	"log"
	"os"
)

type data struct {
	Name string
	Data [][]string
}

var CD data

func Open(fileName string) error {
	cd := &CD
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(f)
	cd.Data, err = csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func GetData() *[][]string {
	return &CD.Data
}
