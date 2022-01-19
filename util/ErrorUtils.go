package util

import (
	"log"
	"net/http"
	"strings"
)

func ThrowNotFoundError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	message := `{"message": "` + err.Error() + `"}`
	w.Write([]byte(message))
	log.Println(err)
}

func ThrowBadRequestError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	message := `{"message": "` + err.Error() + `"}`
	w.Write([]byte(message))
	log.Println(err)
}

func ThrowError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	if strings.Contains(err.Error(), "unable to find pokemon") {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
	message := `{"message": "` + err.Error() + `"}`
	w.Write([]byte(message))
}
