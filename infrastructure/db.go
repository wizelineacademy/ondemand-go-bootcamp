package db

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadData() [][]string {
	file, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records
}
