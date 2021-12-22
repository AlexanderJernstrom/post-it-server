package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string
	PostIts  []PostIt `json:"postits"`
}

func (user *User) HashPassword() error {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		return err
	}

	user.Password = string(hashedBytes)

	return nil
}