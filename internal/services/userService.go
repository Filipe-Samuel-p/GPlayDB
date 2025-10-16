package services

import (
	"gplaydb/internal/models"
	"gplaydb/internal/repositories"
	"time"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (u *UserService) GetAllUsers() ([]models.User, error) {
	return u.Repo.GetAllUsers()
}

func (u *UserService) GetUserById(id string) (*models.User, error) {
	return u.Repo.GetUserById(id)
}

func (u *UserService) InsertUser(newUser *models.User) (models.User, error) {

	newUser.CreatedAt = time.Now()

	return u.Repo.InsertUser(newUser)
}
