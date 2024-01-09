// backend/main.go

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tabintel/gupload/backend/api"
	"github.com/tabintel/gupload/backend/database"
)

func main() {
	// Initialize database
	_, err := database.InitDB()
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Initialize Gin router
	router := gin.Default()

	// Set up API routes
	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/upload", api.HandleUpload)
		apiGroup.GET("/images/:id", api.HandleImage)
	}

	// Serve frontend files
	router.Static("/static", "./frontend/static")
	router.LoadHTMLGlob("frontend/templates/*")

	// Serve the main HTML file
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	// Run the server
	router.Run(":8080")
}
