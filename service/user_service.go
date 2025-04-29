package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gustavsant/go-crud/config"
	"github.com/gustavsant/go-crud/dto"
	"github.com/gustavsant/go-crud/model"
	"github.com/gustavsant/go-crud/security"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterUser(userDto dto.RegisterUserDTO) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"email": userDto.Email}

	var existingUser model.UserModel
	err := config.DB.Collection("users").FindOne(ctx, filter).Decode(&existingUser)
	if err == nil {
		return nil, errors.New("email already in use")
	}

	hashedPassword, err := security.HashPassword(userDto.Password)
	if err != nil {
		return nil, err
	}

	newUser := dto.RegisterUserDTO{
		ID:       userDto.ID,
		Email:    userDto.Email,
		Password: hashedPassword,
	}

	return config.DB.Collection("users").InsertOne(ctx, newUser)
}

func AuthenticateUser(userDto dto.AuthenticateUserDTO) (string, error) {
	email := userDto.Email

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"email": email}
	var user dto.AuthenticateUserDTO

	err := config.DB.Collection("users").FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", fmt.Errorf("no document with email %v in database.", email)
		}
		return "", fmt.Errorf("error trying to retrieve data from database. %v", err)
	}

	if !security.CheckPasswordHash(user.Password, userDto.Password) {
		return "", fmt.Errorf("incorrect password. %v", err)
	}

	tokenString, err := security.GenerateAndSignJWT(user.Email)

	if err != nil {
		return "", fmt.Errorf("failed to generate an token string")
	}

	return tokenString, nil

}

func GetUsers() ([]model.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := config.DB.Collection("users").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []model.UserModel
	err = cursor.All(ctx, &users)

	if err != nil {
		return nil, err
	}

	return users, nil

}
