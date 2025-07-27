package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type NETFLIX struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string             `json:"movie,omitempty" bson:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty" bson:"watched,omitempty"`
}

type Author struct {
	Name    string `json:"name"`
	Website string `json:"website"`
}
