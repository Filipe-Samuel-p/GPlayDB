package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"gplaydb/internal/models"
	"time"

	"github.com/google/uuid"
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
		return nil, fmt.Errorf("Erro ao executar a query: %w", err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt); err != nil {
			return nil, fmt.Errorf("erro ao fazer scan do usuário: %w", err)
		}
		users = append(users, u)
	}

	return users, nil

}

func (r *UserRepository) GetUserById(id uuid.UUID) (*models.User, error) {

	var u models.User
	err := r.DB.QueryRow(
		"SELECT id, name, email, created_at FROM users WHERE id = $1",
		id,
	).Scan(&u.ID, &u.Name, &u.Email, &u.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("usuário com id '%s' não encontrado: %w", id, err)
		}
		return nil, fmt.Errorf("erro ao buscar usuário por id '%s': %w", id, err)
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

func (r *UserRepository) DeleteUserById(id uuid.UUID) error {
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

func (r *UserRepository) UserWithProducts(id uuid.UUID) (*models.User, error) {
	query := `
		SELECT 
			u.id, u.name, u.email, u.created_at,
			p.id, p.name, p.monthly_price
		FROM users u
		INNER JOIN subscriptions s ON s.user_id = u.id
		INNER JOIN products p ON p.id = s.product_id
		WHERE u.id = $1
	`

	rows, err := r.DB.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar query: %w", err)
	}
	defer rows.Close()

	var user models.User
	first := true

	for rows.Next() {
		var (
			p models.Product
		)

		if first {
			err := rows.Scan(
				&user.ID,
				&user.Name,
				&user.Email,
				&user.CreatedAt,
				&p.ID,
				&p.Name,
				&p.Monthly_price,
			)
			if err != nil {
				return nil, fmt.Errorf("erro ao escanear linha: %w", err)
			}
			first = false
		} else {
			var dummyID, dummyName, dummyEmail string
			var dummyCreatedAt time.Time
			err := rows.Scan(
				&dummyID,
				&dummyName,
				&dummyEmail,
				&dummyCreatedAt,
				&p.ID,
				&p.Name,
				&p.Monthly_price,
			)
			if err != nil {
				return nil, fmt.Errorf("erro ao escanear linha: %w", err)
			}
		}

		user.Products = append(user.Products, p)
	}

	if first {
		return nil, fmt.Errorf("usuário não encontrado")
	}

	return &user, nil
}
