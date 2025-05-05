package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavsant/go-crud/controller"
	"github.com/gustavsant/go-crud/middlewares"
)

func MovieRouter(router *gin.Engine) {
	router.GET("/movie/:id", controller.GetMovie)
	router.GET("/movies", controller.GetMovies)

	router.POST("/movies", middlewares.AuthMiddleware(), controller.CreateMovie)
	router.PUT("/movies/:id", middlewares.AuthMiddleware(), controller.UpdateMovie)
	router.DELETE("movies/:id", middlewares.AuthMiddleware(), controller.DeleteMovie)

	router.POST("/register", controller.RegisterUser)
	router.POST("/login", controller.AuthenticateUser)
	router.POST("/logout", controller.LogoutUser)
	router.GET("/me", controller.GetUserInfo)
}
