package services

import (
	"gplaydb/internal/models"
	"gplaydb/internal/repositories"
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
