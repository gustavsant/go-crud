package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gustavsant/go-crud/src/config"
	"github.com/gustavsant/go-crud/src/routes"
)

func main() {

	config.ConnectDatabase()
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Success!",
		})
	})

	routes.MovieRouter(router)

	router.Run(":3220")
}
