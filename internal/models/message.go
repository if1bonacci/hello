package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	Id   primitive.ObjectID `json:"id,omitempty"`
	Body string             `json:"body,omitempty" validate:"required"`
}
