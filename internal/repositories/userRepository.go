package repositories

import (
	"database/sql"

	"gplaydb/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT id, name, email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil

}

func (r *UserRepository) GetUserById(id string) (*models.User, error) {
	var u models.User
	err := r.DB.QueryRow("SELECT id, name, email, created_at FROM users WHERE id = $1", id).
		Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) InsertUser(newUser *models.User) (models.User, error) {
	var u models.User
	query := `
		INSERT INTO users (name, email, created_at)
		VALUES ($1, $2, $3)
		RETURNING id, name, email, created_at
	`

	err := r.DB.QueryRow(query, newUser.Name, newUser.Email, newUser.CreatedAt).
		Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)

	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func (r *UserRepository) DeleteUserById(id string) error {
	query := `DELETE FROM users
			WHERE id = $1`

	_, err := r.DB.Exec(query, id)
	return err
}

func (r *UserRepository) UpdateUser(user *models.User) (*models.User, error) {

	var userAux models.User
	query := ` UPDATE users SET name=$1, email=$2
				WHERE id = $3
				RETURNING id, name, email,created_at`

	err := r.DB.QueryRow(query, user.Name, user.Email, user.ID).
		Scan(&userAux.ID, &userAux.Name, &userAux.Email, &userAux.CreatedAt)
	return &userAux, err

}
