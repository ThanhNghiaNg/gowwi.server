package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Type struct {
	User      bson.ObjectID `json:"user"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
