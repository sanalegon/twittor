package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sanalegon/twittor/bd"
	"github.com/sanalegon/twittor/jwt"
	"github.com/sanalegon/twittor/models"
)

// Todo lo de routers son endpoints
/* Login realiza el login */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t) // procesar el body para ver que tiene dentro(en login seria el email y la pass), luego hacemos el decode dentro de t

	if err != nil {
		http.Error(w, "Usuario y/o Contrasena invalidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	documento, existe := bd.IntentoLogin(t.Email, t.Password)

	if existe == false {
		http.Error(w, "Usuario y/o Contrasena invalidos", 400)
		return
	}

	// jwt son tokens de identidad. que sirven para mantener el acceso de usuario dentro de la app. Asi no tengo que validar todo el tiempo la pass del usuario

	jwtKey, err := jwt.GeneroJWT(documento) // devuelve el token en formato string

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el token correspondiente => "+err.Error(), 400)
		return
	}

	// Modelo formato json. Devolvemos al navegador el jwtKey
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json") // Seteamos el header con formato json
	w.WriteHeader(http.StatusCreated)                  // status 200 al header
	json.NewEncoder(w).Encode(resp)                    // encodificar en mi response writer w y hacerle un encode de resp(uesta)

	// Grabar cookies
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	}) // hay que hacer un puntero hacia la funcion cookie del http
}
