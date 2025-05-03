package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type RegisterUserDTO struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `json:"email" binding:"required,email"`
	Password string             `json:"password" binding:"required,min=8"`
}

type AuthenticateUserDTO struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `json:"email" binding:"required"`
	Password string             `json:"password" binding:"required"`
}
