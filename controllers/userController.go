package controllers

import (
	"go-server/config"
	"go-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.Users
	// Find all users and preload their Educations.
	if err := config.DB.Preload("Educations").Find(&users).Error; err != nil {
		// An error occurred, send a 500 response.
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Successfully retrieved users, send them back as JSON.
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserByID(c *gin.Context) {

}

func CreateUser(c *gin.Context) {

}
