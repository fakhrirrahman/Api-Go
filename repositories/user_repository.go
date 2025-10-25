package repositories

import (
	"fmt"
	"goApi/interfaces"
	"goApi/models"
)

type userRepository struct {
	data []models.User
}

func NewUserRepository() interfaces.UserInterface {
	return &userRepository{
		data: []models.User{},
	}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	return r.data, nil
}

func (r *userRepository) GetByID(id int) (models.User, error) {
	for _, user := range r.data {
		if user.ID == id {
			return user, nil
		}
	}
	return models.User{}, nil
}

func (r *userRepository) Create(user models.User) error {
	for _, u := range r.data {
		if u.Email == user.Email {
			return fmt.Errorf("user dengan email %s sudah ada", user.Email)
		}
	}

	user.ID = len(r.data) + 1

	r.data = append(r.data, user)
	return nil
}
