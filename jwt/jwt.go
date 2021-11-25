package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sanalegon/twittor/models"
)

/* GeneroJWT genera el encriptado con JWT */
func GeneroJWT(t models.Usuario) (string, error) {

	// Creamos una clave privada. JWT trabaja con un slice de bytes, no con strings
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")

	// Claims son los privilegios. Creamos la lista de privilegios
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(), // El formato unix convierte el formato fecha en un long. Es ilegible pero es muy rapido y liviano
	}

	// Aqui elegimos el algoritmo que ha de tener en cuenta para encriptar y hacer todo el calculo. Esta es info para el header
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	// ahora al token lo firmamos(siendo esto la ultima parte del token)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
