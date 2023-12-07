package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pseudonative/web_app_kube/internal/db/database"
	"github.com/pseudonative/web_app_kube/internal/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal("Could not connect to the database")
	}
	defer db.Close()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello World again"})
	})

	router.GET("/users/:id", handlers.GetUserHandler(db))
	router.GET("/users", handlers.GetAllUsersHandler(db))
	router.POST("/users", handlers.CreateUserHandler(db))
	router.PATCH("/users/:id", handlers.UpdateUserHandler(db))
	router.DELETE("/users/:id", handlers.DeleteUserHandler(db))

	router.Run(":8080")
}
