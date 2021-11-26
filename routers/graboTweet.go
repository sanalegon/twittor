package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sanalegon/twittor/bd"
	"github.com/sanalegon/twittor/models"
)

/* GraboTweet permite grabar el tweet en la base de datos */
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje) // Decode es hacer una cadena de ejecucion que toma el puntero de nuestro mensaje

	// Creo una variable de tipo models grabo tweet y ahi grabo toda la info
	registro := models.GraboTweet{
		UserID:  IDUsuario, //variable global que usare ahora. esta sale cuando se establecio el token
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el regisro, reintente nuevamente "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
