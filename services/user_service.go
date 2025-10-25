package services

import (
	"goApi/interfaces"
	"goApi/models"
)

type UserService struct {
	repo interfaces.UserInterface
}

func NewUserService(r interfaces.UserInterface) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetUserrByID(id int) (models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) AddUser(u models.User) error {
	return s.repo.Create(u)
}
