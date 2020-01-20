package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"../data"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func GetPeople(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(data.People)
}

func GetPerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range data.People {
		if item.ID == bson.ObjectIdHex(params["id"]) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&data.Person{})
}

func UpdatePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range data.People {
		if item.ID == bson.ObjectIdHex(params["id"]) {
			var newPerson data.Person
			err := json.NewDecoder(req.Body).Decode(&newPerson)

			if err != nil {
				log.Println("Error while updating Person.", err)
				return
			}

			if newPerson.FirstName != "" {
				data.People[index].FirstName = newPerson.FirstName
			}
			if newPerson.LastName != "" {
				data.People[index].LastName = newPerson.LastName
			}
			if newPerson.Job != "" {
				data.People[index].Job = newPerson.Job
			}

			json.NewEncoder(w).Encode(data.People[index])
			return
		}
	}
	json.NewEncoder(w).Encode(&data.Person{})
}

func CreatePerson(w http.ResponseWriter, req *http.Request) {
	var person data.Person
	err := json.NewDecoder(req.Body).Decode(&person)

	if err != nil {
		log.Println("Error while creating new Person.", err)
		return
	}

	obj_id := bson.NewObjectId()
	person.ID = obj_id

	data.People = append(data.People, person)
	json.NewEncoder(w).Encode(data.People)
}

func DeletePerson(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range data.People {
		if item.ID == bson.ObjectIdHex(params["id"]) {
			data.People = append(data.People[:index], data.People[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(data.People)
}
