package middlew

import (
	"net/http"

	"github.com/sanalegon/twittor/routers"
)

//Los middleware reciben y devuelven lo mismo!!!
/* ValidoJWT permite validar el JWT que nos viene en la peticion*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization")) // Aqui accedemos de postman a la variabble authorization. Se uso de ejemplo para econtrarla LeoTweets y de ahi se encuentra en headers esta variable

		if err != nil {
			http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
			return
		}
		// Si todo sale bien, pasamos los dos objetos al siguiente eslabon de la cadena
		next.ServeHTTP(w, r)
	}
}
