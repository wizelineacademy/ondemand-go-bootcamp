package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/gorilla/mux"
)

type Pokemon struct {
	ID         string `csv:"ID"`
	Name       string `csv:"Name"`
	Type_1     string `csv:"Type_1"`
	Type_2     string `csv:"Type_2"`
	Total      int    `csv:"Total"`
	HP         int    `csv:"HP"`
	Attack     int    `csv:"Attack"`
	Defense    int    `csv:"Defense"`
	Sp_Atk     int    `csv:"Sp_Atk"`
	Sp_Def     int    `csv:"Sp_Def"`
	Speed      int    `csv:"Speed"`
	Generation int    `csv:"Generation"`
	Legendary  bool   `csv:"Legendary"`
}

func getAll(w http.ResponseWriter, req *http.Request) {
	content, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}
	pokemons := []*Pokemon{}

	if err := gocsv.UnmarshalFile(content, &pokemons); err != nil {
		panic(err)
	}
	for _, pokemon := range pokemons {
		fmt.Fprintf(w, "Id: %v, Name: %v\n", pokemon.ID, pokemon.Name)
	}
}

func get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		fmt.Fprint(w, "id is missing in parameters")
		return
	}

	in, err := os.Open("data.csv")
	if err != nil {
		log.Fatal((err))
	}
	pokemons := []*Pokemon{}

	if err := gocsv.UnmarshalFile(in, &pokemons); err != nil {
		panic(err)
	}

	for _, pokemon := range pokemons {
		if pokemon.ID == id {
			fmt.Fprintf(w, "%+v\n", pokemon)
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/pokemon", getAll)
	r.HandleFunc("/pokemon/{id}", get)

	http.ListenAndServe("localhost:8090", r)
}
