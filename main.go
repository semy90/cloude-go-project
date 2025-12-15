package main

import (
	"cloudego/rest"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	app := rest.New(r)
	app.Start()
}
