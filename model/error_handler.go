package model

import (
	"fmt"
	"net/http"
)

type ErrorHandler struct {
	SystemCode int    `json:"systemCode"`
	StatusCode int    `json:"statusCode"`
	Title      string `json:"title"`
	Message    string `json:"message"`
}

func NewOpenFileError(err string) *ErrorHandler {
	return &ErrorHandler{
		Title:      "There was a problem with the file",
		Message:    err,
		SystemCode: 1000,
		StatusCode: http.StatusBadRequest,
	}
}

func NewAccesingCSVFileError(err string) *ErrorHandler {
	return &ErrorHandler{
		Title:      "There was a problem accesing to csv file",
		Message:    err,
		SystemCode: 1001,
		StatusCode: http.StatusBadRequest,
	}
}

func NewUnmarshalFileError(err string) *ErrorHandler {
	return &ErrorHandler{
		Title:      "There was a problem unmarshaling file",
		Message:    err,
		SystemCode: 1002,
		StatusCode: http.StatusBadRequest,
	}
}

func NewUnmarshalResponseBodyExternalService(err string) *ErrorHandler {
	return &ErrorHandler{
		Title:      "There was a problem reading responsebody from external service",
		Message:    err,
		SystemCode: 1003,
		StatusCode: http.StatusBadRequest,
	}
}

func NewURLParameterDoesNotFound(parameter string) *ErrorHandler {
	return &ErrorHandler{
		Title:      "The parameter does not found",
		Message:    fmt.Sprintf("The parameter: %s is not present", parameter),
		SystemCode: 1004,
		StatusCode: http.StatusUnprocessableEntity,
	}
}

func NewNotFoundPokemonsError() *ErrorHandler {
	return &ErrorHandler{
		Title:      "The pokemons do not found",
		Message:    "The pokemons do not found",
		SystemCode: 2000,
		StatusCode: http.StatusNotFound,
	}
}

func NewNotFoundPokemonError(id int) *ErrorHandler {
	return &ErrorHandler{
		Title:      "The pokemon does not found",
		Message:    fmt.Sprintf("The pokemon with id: %d does not found", id),
		SystemCode: 2001,
		StatusCode: http.StatusNotFound,
	}
}

func NewPokemonAPIIsNotReached(err string) *ErrorHandler {
	return &ErrorHandler{
		Title:      "The external service is not reached",
		Message:    err,
		SystemCode: 3000,
		StatusCode: http.StatusBadRequest,
	}
}
func NewGetPokemonFromAPINotFoundError(id int) *ErrorHandler {
	return &ErrorHandler{
		Title:      "The Pokemon does not found in the external service",
		Message:    fmt.Sprintf("The pokemon with id: %d does not found", id),
		SystemCode: 3001,
		StatusCode: http.StatusNotFound,
	}
}
