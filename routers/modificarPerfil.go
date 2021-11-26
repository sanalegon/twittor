package routers

import (
	"encoding/json"
	"net/http"

	"github.com/sanalegon/twittor/bd"
	"github.com/sanalegon/twittor/models"
)

/* ModficarPerfil modifica el perfil del usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t) // hacemos un decode en la posicion de memoria t que es decodificar un modelo usuario al actual usuario a modificar

	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	// Al importar routers, tengo acceso a la variable global IDUsuario. Asi evito llamar validar token todo el rato solo para conseguir el ID usuario
	status, err = bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated) // envio que el registro fue modificado. es solo un estado de creado
}
