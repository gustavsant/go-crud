package service

import (
	"context"
	"testing"
	"time"

	"github.com/gustavsant/go-crud/config"
	"github.com/gustavsant/go-crud/model"
	"go.mongodb.org/mongo-driver/bson"
)

func TestGetMovies(t *testing.T) {

	config.ConnectDatabase()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	movie := model.Movie{
		Title:       "test_title",
		Description: "desc_test",
		Rating:      3.3,
		Cover:       "https://www.example.com/image.jpeg",
	}

	res, err := config.DB.Collection("movies").InsertOne(ctx, movie)

	if err != nil {
		t.Fatalf("can't add movie to the db. %v", err)
	}

	movies, err := GetMovies()
	if err != nil {
		t.Fatalf("an nil was waited, but it returned error: %v", err)
	}

	found := false

	for _, movie := range movies {
		if movie.Title == "test_title" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Movie not found on the database.")
	}
	if res != nil {
		_, _ = config.DB.Collection("movies").DeleteOne(ctx, bson.M{"_id": res.InsertedID})
	}

}
