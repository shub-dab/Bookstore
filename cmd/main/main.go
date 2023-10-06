package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shub-dab/bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

