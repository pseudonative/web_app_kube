package queries

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pseudonative/web_app_kube/internal/db/sql/queries"
	"github.com/pseudonative/web_app_kube/internal/models"
)

func CreateUser(db *sql.DB, user models.User) (int, error) {
	var userID int
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id;`
	if err := db.QueryRow(query, user.Name, user.Email).Scan(&userID); err != nil {
		return 0, err
	}
	return userID, nil
}

func GetUsers(db *sql.DB) ([]models.User, error) {
	var users []models.User
	query := `SELECT id, name, email, created_at FROM users;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUser(db *sql.DB, id int) (models.User, error) {
	var user models.User
	query := `SELECT id