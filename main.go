package main

import (
	"net/http"

	router "pas.com/v1/server"
)

func main() {

	http.ListenAndServe(":8090", router.Router())
}
