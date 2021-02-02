package db

import (
	"github.com/santyrh70/twittar/models"
	"golang.org/x/crypto/bcrypt"
)

func IntentoLogin(email string, password string) (models.Ususario, bool) {
	usu, encontrado, _ := ChequeoExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}

	passwordDBytes := []byte(password)
	passwordDB := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordDBytes)
	if err != nil {
		return usu, false
	}

	return usu, true
}
