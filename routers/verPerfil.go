package routers

import (
	"encoding/json"
	"net/http"

	"github.com/sanalegon/twittor/bd"
)

/* VerPerfil permite extraer los valores del Perfil */
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	// Extraer del body los parametros que vinieron
	ID := r.URL.Query().Get("id") // Lo extraemos del objeto request
	if len(ID) < 1 {
		// si cero, es que no lo encontro
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		// perfil no encontrado
		http.Error(w, "Ocurrio un error al intentar buscar el registro "+err.Error(), 400)
		return
	}

	// Seteamos para decirle que estamos seteando un json
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil) // Encodeamos lo que tiene que ver con el json
}
