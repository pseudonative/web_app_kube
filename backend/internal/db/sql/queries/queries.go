package queries

import (
	"database/sql"

	"github.com/pseudonative/web_app_kube/internal/models"
)

// GetUserByID retrieves a user by their ID from the database.
func GetUserByID(db *sql.DB, id int) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetAllUsers retrieves all users from the database.
func GetAllUsers(db *sql.DB) ([]models.User, error) {
	var users []models.User
	query := `SELECT id, name, email, created_at FROM users`
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

	return users, rows.Err()
}

// CreateUser inserts a new user into the database.
func CreateUser(db *sql.DB, user models.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	return db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
}

// UpdateUser updates an existing user's information in the database.
func UpdateUser(db *sql.DB, id int, name, email string) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err := db.Exec(query, name, email, id)
	return err
}

// DeleteUser removes a user from the database by their ID.
func DeleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}
