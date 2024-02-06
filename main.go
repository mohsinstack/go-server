package main

import (
	"go-server/config"
	"go-server/models"
	"go-server/routes"

	"github.com/gin-contrib/cors"
)

func main() {
	// Establish a connection to our database
	config.ConnectDatabase()

	// Migrate the schema for Users and Education Details
	config.DB.AutoMigrate(&models.Users{}, &models.Education{})

	// Set up our web server's routes
	r := routes.SetupRouter()
	// CORS middleware to allows or blocks browsers from accessing our
	// server from different places on the internet.
	r.Use(cors.Default())
	// Start server and listen port 8080
	r.Run(":8080")
}
