package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// creating object to insert it in db
type Match struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Teams string             `json:"teams,omitempty"`
	Date  string             `json:"date"`
	Was   bool               `json:"was,omitempty"`
}
