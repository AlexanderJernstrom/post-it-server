package service

import (
	"server/database"
	"server/models"
)




func CreateUser(user models.User) (models.User, error) {
	if err := user.HashPassword(); err != nil {
		return models.User{}, err
	}

	if result := database.Client.Create(&user).First(&user); result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func UserWithEmailExists(email string) (bool, error) {
	var user []models.User

	if result := database.Client.Where("email = ?", email).Find(&user); result.Error != nil {
		return false, result.Error
	}

	if len(user) > 0 {
		return true, nil
	}

	return false, nil
}


func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	if db := database.Client.Where("email = ?", email).First(&user); db.Error != nil {
		return models.User{}, db.Error
	}

	return user, nil


}