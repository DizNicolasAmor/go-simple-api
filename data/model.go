package data

import "gopkg.in/mgo.v2/bson"

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty" json:"id"`
	FirstName string        `json:"firstname"`
	LastName  string        `json:"lastname"`
	Job       string        `json:"job"`
}
