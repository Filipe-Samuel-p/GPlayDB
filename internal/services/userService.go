package services

import (
	"gplaydb/internal/models"
	"gplaydb/internal/repositories"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) GetUserById(id uuid.UUID) (*models.User, error) {
	return s.Repo.GetUserById(id)
}

func (s *UserService) InsertUser(newUser *models.User) (models.User, error) {

	newUser.CreatedAt = time.Now()

	return s.Repo.InsertUser(newUser)
}

func (s *UserService) DeleteUserById(id uuid.UUID) error {
	return s.Repo.DeleteUserById(id)
}

func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {

	originalUser, err := s.GetUserById(user.ID)
	if err != nil {
		return nil, err
	}

	if user.Name != "" {
		originalUser.Name = user.Name
	}
	if user.Email != "" {
		originalUser.Email = user.Email
	}

	userUpdate, err := s.Repo.UpdateUser(originalUser)

	return userUpdate, err
}

func (s *UserService) UserWithProducts(id uuid.UUID) (*models.User, error) {
	return s.Repo.UserWithProducts(id)
}
