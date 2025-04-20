package service

import (
	"context"
	"errors"
	"time"

	"github.com/gustavsant/go-crud/src/config"
	"github.com/gustavsant/go-crud/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMovie(movie model.Movie) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return config.DB.Collection("movies").InsertOne(ctx, movie)

}

func GetMovies() ([]model.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.DB.Collection("movies").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var movies []model.Movie
	if err = cursor.All(ctx, &movies); err != nil {
		return nil, err
	}

	return movies, nil

}

func UpdateMovie(id string, updates *model.Movie) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	updateData := bson.M{}

	if updates.Title != "" {
		updateData["title"] = updates.Title
	}

	if updates.Rating != 0 {
		updateData["rating"] = updates.Rating

	}

	if updates.Description != "" {
		updateData["description"] = updates.Description

	}

	if updates.Cover != "" {
		updateData["cover"] = updates.Cover

	}

	if len(updateData) == 0 {
		return errors.New("no valid field for changes")
	}

	update := bson.M{"$set": updateData}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = config.DB.Collection("movies").UpdateOne(ctx, filter, update)
	return err
}

func DeleteMovie(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := config.DB.Collection("movies").DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no movie found with that id")
	}

	return nil
}
