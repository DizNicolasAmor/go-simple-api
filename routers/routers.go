package routers

import (
	"../handlers"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/persons", handlers.CreatePerson).Methods("POST")
	router.HandleFunc("/persons", handlers.GetPersons).Methods("GET")
	router.HandleFunc("/persons/{id}", handlers.GetPersonByID).Methods("GET")
	router.HandleFunc("/persons/{id}", handlers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/persons/{id}", handlers.DeletePerson).Methods("DELETE")

	return router
}
