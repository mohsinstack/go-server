package controllers

import (
	"go-server/config"
	"go-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	// Extract the ID from the request URL.
	id := c.Param("id")

	var user models.Users
	// Attempt to find the item by ID and preload its Educations.
	if err := config.DB.Preload("Educations").First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Item not found, send a 404 response.
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		} else {
			// An unexpected error occurred, send a 500 response.
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Successfully found the item, send it back as JSON.
	c.JSON(http.StatusOK, gin.H{"user": user})

}

func CreateUser(c *gin.Context) {
	// Define a struct to match the expected input.
	var userInput struct {
		Firstname  string             `json:"firstname"`
		Lastname   string             `json:"lastname"`
		Email      string             `json:"email"`
		Educations []models.Education `json:"educations"`
	}
	// Bind the incoming JSON to the struct.
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new User from the input.
	user := models.Users{
		Firstname:  userInput.Firstname,
		Lastname:   userInput.Lastname,
		Email:      userInput.Email,
		Educations: userInput.Educations,
	}

	// Save the new item to the database.
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Item creation was successful, send back the created item.
	c.JSON(http.StatusCreated, gin.H{"user": user})
}
