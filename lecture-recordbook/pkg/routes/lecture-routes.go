package routes

import (
	"github.com/akhil/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterLecStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/lec/", controllers.CreateLec).Methods("POST")
	router.HandleFunc("/lec/", controllers.GetLec).Methods("GET")
	router.HandleFunc("/lec/{lecId}", controllers.GetLecById).Methods("GET")
	router.HandleFunc("/lec/{lecId}", controllers.UpdateLec).Methods("PUT")
	router.HandleFunc("/lec/{lecId}", controllers.DeleteLec).Methods("DELETE")

}
