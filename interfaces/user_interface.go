package interfaces

import "goApi/models"

type UserInterface interface {
	GetAll() ([]models.User, error)
	GetByID(id int) (models.User, error)
	Create(user models.User) error
}
