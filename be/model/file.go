package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type File struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name" validate:"required,name"`
	Email string             `json:"email" bson:"email" validate:"required,email"`
}
