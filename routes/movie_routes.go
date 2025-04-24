package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavsant/go-crud/controller"
)

func MovieRouter(router *gin.Engine) {
	router.POST("/movies", controller.CreateMovie)
	router.GET("/movies", controller.GetMovies)
	router.PUT("/movies/:id", controller.UpdateMovie)
	router.DELETE("movies/:id", controller.DeleteMovie)
	router.GET("/movie/:id", controller.GetMovie)
}
