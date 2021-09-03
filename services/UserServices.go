package services

import (
	"errors"
	"restapidemo/models"

	"gorm.io/gorm"
)

type User models.User

type UserService struct{}

func (userService *UserService) GetUserById(userId string) (User, error) {

	var user User

	result := db.Find(&user, userId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("User Id not found")
	}

	return user, nil
}

func (userService *UserService) CreateUser(userRequest *models.User) error {
	err := db.Model(User{}).Create(userRequest)

	if err != nil {
		return errors.New("Failed to create User")
	}

	return nil
}
