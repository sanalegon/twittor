package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sanalegon/twittor/bd"
)

/* LeoTweetsSeguidores lee los tweets de todos nuestros seguidores */
func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina como entero mayor a cero", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeoTweetsSegudores(IDUsuario, pagina)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
