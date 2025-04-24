package service

import (
	"context"
	"errors"
	"time"

	"github.com/gustavsant/go-crud/config"
	"github.com/gustavsant/go-crud/dto"
	"github.com/gustavsant/go-crud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateMovie(movieDTO dto.CreateMovieDTO) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return config.DB.Collection("movies").InsertOne(ctx, movieDTO)

}

func GetMovies() ([]dto.MovieResponseDTO, error) {
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

	var movieDTOs []dto.MovieResponseDTO

	for _, movie := range movies {
		movieDTO := dto.MovieResponseDTO{
			ID:          movie.ID.Hex(),
			Title:       movie.Title,
			Description: movie.Description,
			Rating:      movie.Rating,
			Cover:       movie.Cover,
		}

		movieDTOs = append(movieDTOs, movieDTO)
	}

	return movieDTOs, nil

}

func GetMovie(id string) (dto.MovieResponseDTO, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dto.MovieResponseDTO{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectId}

	var movie dto.MovieResponseDTO
	err = config.DB.Collection("movies").FindOne(ctx, filter).Decode(&movie)

	if err != nil {
		return dto.MovieResponseDTO{}, err
	}

	return movie, nil

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
