package bd

import (
	"github.com/sanalegon/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/* IntentoLogin realiza el chequeo de login a la BD */
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)

	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	// contras no son iguales si es distinto de nil
	if err != nil {
		return usu, false
	}

	return usu, true
}
