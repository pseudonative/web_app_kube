package queries

import (
	"database/sql"
	"log"

	"github.com/pseudonative/web_app_kube/pkg/models"
)

func CreateUser(db *sql.DB, user models.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id;`
	_, err := db.Exec(query, user.Name, user.Email)
	if err != nil {
		log.Printf("Error creating new user: %v", err)
		return err
	}
	return db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
}

func GetUser(db *sql.DB, id int) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1;`
	row := db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		log.Printf("Error getting user: %v", err)
		return nil, err
	}
	return user, nil
}

func UpdateUsers(db *sql.DB, id int, name, email string) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3;`
	_, err := db.Exec(query, name, email, id)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

func DeleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1;`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error Deleting user: %v", err)
		return err
	}
	return nil
}
