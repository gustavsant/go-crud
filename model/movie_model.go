package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Rating      float32            `bson:"rating" json:"rating"`
	Cover       string             `bson:"cover" json:"cover"`
}
