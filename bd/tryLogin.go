package bd

import (
	"github.com/haguirre92/twitter1/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, pass string) (models.User, bool) {
	usu, find, _ := CheckExistUser(email)
	if find == false {
		return usu, false
	}

	passwordBytes := []byte(pass)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true
}
