package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pseudonative/web_app_kube/pkg/db/database"
	"github.com/pseudonative/web_app_kube/pkg/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.ConnectToDB()
	defer db.Close() // Ensure this is the last defer for db so it closes as the program exits.

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World again"})
	})

	// Register your routes and handlers
	router.GET("/users/:id", handlers.GetUser(db)) // This should probably be for getting a single user
	router.GET("/users", handlers.GetUsers(db))    // This should probably be for listing all users
	router.POST("/users", handlers.CreateUser(db))
	router.PATCH("/users/:id", handlers.UpdateUser(db))
	router.DELETE("/users/:id", handlers.DeleteUser(db))

	router.Run(":8080") // Start the server with this single call.
}
