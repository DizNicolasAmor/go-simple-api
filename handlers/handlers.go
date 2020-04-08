package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"../data"
	"../models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// CreatePerson creates a new person with users submitted data
func CreatePerson(w http.ResponseWriter, req *http.Request) {
	var person models.Person
	err := json.NewDecoder(req.Body).Decode(&person)

	if err != nil {
		log.Println("Bad request.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	person.ID = bson.NewObjectId()

	data.Persons = append(data.Persons, person)

	j, err := json.Marshal(person)
	if err != nil {
		log.Println("Error marshalling json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// GetPersons returns an array with all persons saved in data package
func GetPersons(w http.ResponseWriter, req *http.Request) {
	j, err := json.Marshal(data.Persons)
	if err != nil {
		log.Println("Error marshalling json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// GetPersonByID gets person by ID if exist
func GetPersonByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if bson.IsObjectIdHex(id) {
		for _, item := range data.Persons {
			if item.ID == bson.ObjectIdHex(id) {
				j, err := json.Marshal(item)
				if err != nil {
					log.Println("Error marshalling json")
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(j)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

// UpdatePerson finds person by ID and if exist, updates its data
func UpdatePerson(w http.ResponseWriter, req *http.Request) {
	var body models.Person
	err := json.NewDecoder(req.Body).Decode(&body)

	if err != nil {
		log.Println("Bad request.")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(req)
	id := vars["id"]

	if bson.IsObjectIdHex(id) {
		for index, item := range data.Persons {
			if item.ID == bson.ObjectIdHex(id) {
				data.Persons[index].FirstName = body.FirstName
				data.Persons[index].LastName = body.LastName
				data.Persons[index].Job = body.Job

				j, err := json.Marshal(data.Persons[index])
				if err != nil {
					log.Println("Error marshalling json")
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(j)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}

// DeletePerson searchs person by ID and deletes it if exist
func DeletePerson(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if bson.IsObjectIdHex(id) {
		for index, item := range data.Persons {
			if item.ID == bson.ObjectIdHex(id) {
				data.Persons = append(data.Persons[:index], data.Persons[index+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
}
