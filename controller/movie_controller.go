package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavsant/go-crud/dto"
	"github.com/gustavsant/go-crud/model"
	"github.com/gustavsant/go-crud/service"
)

func CreateMovie(ctx *gin.Context) {
	var movie dto.CreateMovieDTO
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := service.CreateMovie(movie)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func GetMovies(ctx *gin.Context) {
	movies, err := service.GetMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, movies)
}

func GetMovie(ctx *gin.Context) {
	id := ctx.Param("id")

	result, err := service.GetMovie(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, result)
}

func UpdateMovie(ctx *gin.Context) {
	id := ctx.Param("id")

	var movieUpdates model.Movie

	if err := ctx.ShouldBindJSON(&movieUpdates); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := service.UpdateMovie(id, &movieUpdates)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "movie of id:" + id + " changed successfully",
	})
}

func DeleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")

	err := service.DeleteMovie(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "the movie with id: " + id + " was removed successfully",
	})
}
