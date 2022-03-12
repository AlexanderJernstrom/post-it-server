package models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)


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

func (user *User) CreateToken() (token string, err error) {
	tokenClaims := jwt.MapClaims{}
	tokenClaims["authorized"] = true
	tokenClaims["user_id"] = user.ID

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	token, err = at.SignedString([]byte("SECRET"))
	if err != nil {
		return "", err
	}
	
	return token, nil
}