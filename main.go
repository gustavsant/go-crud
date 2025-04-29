package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gustavsant/go-crud/config"
	"github.com/gustavsant/go-crud/routes"
)

func main() {

	config.ConnectDatabase()
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Pong!",
		})
	})

	routes.MovieRouter(router)

	router.Run(":3220")
}
