package controllers

import (
	"github.com/LuanFreitasRibeiro/webapi-mvc-golang.git/database"
	"github.com/LuanFreitasRibeiro/webapi-mvc-golang.git/models"
	"github.com/LuanFreitasRibeiro/webapi-mvc-golang.git/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	user.Password = services.Sha256Encoder(user.Password)

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

func GetUsers(c *gin.Context) {
	db := database.GetDatabase()
	var users []models.User
	err := db.Find(&users).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find product by id: " + err.Error(),
		})

		return
	}

	c.JSON(200, users)
}
