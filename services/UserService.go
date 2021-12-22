package service

import (
	"server/database"
	"server/models"
)




func CreateUser(user models.User) (models.User, error) {
	if err := user.HashPassword(); err != nil {
		return models.User{}, err
	}

	if result := database.Client.Create(user); result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}