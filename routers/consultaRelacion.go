package routers

import (
	"encoding/json"
	"net/http"

	"github.com/sanalegon/twittor/bd"
	"github.com/sanalegon/twittor/models"
)

func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)

	// No mando el error porque en el frontend, no les interesa en esta funcion eso. Solo quieren saber si hay o no relacion. Por eso la variable status se setea segun resp(uesta)

	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
