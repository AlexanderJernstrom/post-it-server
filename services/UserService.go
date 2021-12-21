package service

import (
	"server/database"
	"server/models"

	"golang.org/x/crypto/bcrypt"
)



func hashPassword(user *models.User) error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		return err
	}

	user.Password = string(hashedBytes)

	return nil
}

func CreateUser(user models.User) (models.User, error) {
	if err := hashPassword(&user); err != nil {
		return models.User{}, err
	}

	if result := database.Client.Create(user); result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}