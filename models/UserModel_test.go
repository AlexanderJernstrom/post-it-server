package models

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	testUser := User{Name: "test", Password: "test", Email: "Hello"}
	testUser.HashPassword()
	if err := bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte("test")); err != nil {
		t.Error("Hashing password failed")
	}
}