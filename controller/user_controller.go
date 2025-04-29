package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gustavsant/go-crud/dto"
	"github.com/gustavsant/go-crud/security"
	"github.com/gustavsant/go-crud/service"
)

func RegisterUser(c *gin.Context) {
	var userDto dto.RegisterUserDTO

	err := c.ShouldBindJSON(&userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error creating a new user.",
		})
		return
	}

	result, err := service.RegisterUser(userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetUsers(c *gin.Context) {
	users, err := service.GetUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get users",
		})
	}

	c.JSON(http.StatusOK, users)
}

func AuthenticateUser(c *gin.Context) {
	var userDTO dto.AuthenticateUserDTO

	err := c.ShouldBindJSON(&userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid login request.",
		})
		return
	}

	tokenString, err := service.AuthenticateUser(userDTO)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, false)

	c.JSON(http.StatusOK, tokenString)

}

func GetUserInfo(c *gin.Context) {
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid cookie ",
		})
		return
	}

	claims, err := security.ValidateJWT(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email": claims.UserEmail,
	})
}

func LogoutUser(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "successful logout",
	})

}
