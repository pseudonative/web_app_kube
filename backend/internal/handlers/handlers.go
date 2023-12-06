package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/pseudonative/web_app_kube/pkg/db/sql/queries"
	"github.com/pseudonative/web_app_kube/pkg/models"
)

func CreateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		// Bind the JSON to the user struct
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Perform the SQL insert operation
		query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
		err := db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
		if err != nil {
			// Here you need to check the type of error returned.
			// For example, if it's a unique constraint failure, send a specific error message.
			if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
				c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
				return
			}
			// For any other type of error, log it and send a generic error response.
			log.Printf("Error creating new user: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		// If successful, send back the created user object
		c.JSON(http.StatusCreated, gin.H{"user": user})
	}
}

func GetUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
			return
		}

		user, err := queries.GetUser(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func GetUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func UpdateUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, _ := strconv.Atoi(c.Param("id")) // handle error appropriately

		// Update user in the database
		err := queries.UpdateUsers(db, id, user.Name, user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User updated Successfully", "user": user})
	}
}

func DeleteUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id")) // handle error appropriately

		// Delete user from the database
		err := queries.DeleteUser(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted Successfully"})
	}
}
