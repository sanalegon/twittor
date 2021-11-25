package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sanalegon/twittor/bd"
	"github.com/sanalegon/twittor/models"
)

// Variables a exportar a todos mis paquetes de ruta
/* Email valor de Email usado en todos los EndPoints */
var Email string

/* IDUsuario es el ID devuelto del modelo, que se usara en todos los EndPoints */
var IDUsuario string

//Importante: Cuando una funcion devuelve muchos parametros, y uno de ellos es un error, este debe ir siempre al final
/* ProcesoToken procesa token para extraer sus valores*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{} // Lo ponemos como un puntero. Ya que asi los procesa JWT. La estrucutra a chequear tiene que ser un puntero, no acepta otra cosa

	splitToken := strings.Split(tk, "Bearer") // El token arranca con la palbra Bearer. Esta funcion split me devolvera un vector con posicion uno la palabra bearer y posicion 2 el token en si

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token de invalido")
	}

	tk = strings.TrimSpace(splitToken[1]) // borramos los espacios en blanco del token

	// claims es donde guardo mi token
	// interface es muy usado en json para convertir una estructura a jsonn
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		// convertimos miClave en un objeto json(tipo interface)
		return miClave, nil
	})

	if err == nil {
		// primero es ignorado porque ya tenemos claim con la info del usuario
		//_, encontrado, ID := bd.ChequeoYaExisteUsuario(claims.Email)
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)

		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
