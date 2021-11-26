package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sanalegon/twittor/bd"
)

/* LeoTweet leo los tweets */
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	// extreamos la id de la url y no del body, porque esto sera un metdo get (?
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	// primero me aseguro de que vino algo
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	//Ahora intento hacer la conversion
	// todo lo que recibimos por url, esta en formato string, pero aqui toca convertirlo a entero
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina")) // no se convirtio previamente, ya que si quiero convertir a Atoi(to integer)

	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina con un valor mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pagina) // convertimos a pagina de tipo int a int64. En bson se pagina con int64
	respuesta, correcto := bd.LeoTweets(ID, pag)

	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json") // dedcimos de que tipo sera el header
	w.WriteHeader(http.StatusCreated)                  // decimos en el header que fue satisfactorio
	json.NewEncoder(w).Encode(respuesta)
}
