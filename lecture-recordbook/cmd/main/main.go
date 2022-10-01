package main

import (
	"log"
	"net/http"

	"github.com/akhil/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterLecStoreRoutes(r)
	http.Handle("/", r)
	err := http.ListenAndServe("localhost:9010", r)
	if err != nil {
		log.Println("Server is started at localhost: 9010 successfully")
	} else {
		log.Fatalln(err)
	}
}
