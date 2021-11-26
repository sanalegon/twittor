package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sanalegon/twittor/middlew"
	"github.com/sanalegon/twittor/routers"
)

// Funcion a ejecutar al ser llamada la API
/* Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores() {
	// con mux capturamos el http y daremos al response writer y ... tipo si en el header hay info. Devuelve un token/estado
	router := mux.NewRouter()

	// Rutas(Endpoints):
	// si endpoint es registro y llego por post, ejecutara el middleware. Chequeamos que la bd es ok, de ser asi, devolver el control al router de registro
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST") // endpoint del local post. Con chequeo miramos si el return de esta funcion es correcto, hara un (routers.) registro
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET") // Aqui hay un middleware para validar el token
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")

	PORT := os.Getenv("PORT") // si no hay una variable de entorno llamada port, la configuraremos
	if PORT == "" {
		PORT = "8080"
	}

	// cors son los permisos que le doy a mi api para que sea accesible everywhere
	handler := cors.AllowAll().Handler(router) // damos permiso a cualquiera. En el futuro es mejor usar otros permisos
	// escuchamos el puerto, seteamos el puerto ":"+port y le pasamos el handler con los permisos
	log.Fatal(http.ListenAndServe(":"+PORT, handler)) // Ponemos el servidor a escuchar y ver todos los llamados de peticiones
}
