package middlew

import (
	"net/http"

	"github.com/sanalegon/twittor/bd"
)

// Un middleware siempre que recive algo, devuelve lo mismo
// devolvemos una funcion tipo http
/* ChequeoBD es el middleware que me permite conocer el estado de la bd */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	// W y R son los nombres de lo que retorno. * es para decir que es un puntero
	// Ejm: return true ahora lo que hacemos es un retorno de una funcion anonima
	return func(w http.ResponseWriter, r *http.Request) {
		// Si hubo un error en la conexion:
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de Datos", 500) // w es el objeto que he creado que lo usare para trabajar. 500 es el codigo de estado
			return                                                      // Matamos toda la cadena de llamadas
		}
		// next tambien podria ser llamado proximo endpoint. Es solo un nombre
		next.ServeHTTP(w, r) // Devolvemos al proximo endpoint, todos los objetos de w y r
	}
}
