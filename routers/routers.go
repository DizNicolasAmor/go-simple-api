package routers

import (
	"../handlers"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/people", handlers.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/people", handlers.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", handlers.DeletePerson).Methods("DELETE")

	return router
}
