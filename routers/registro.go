package routers

import (
	"encoding/json"
	"net/http"

	"github.com/sanalegon/twittor/bd"
	"github.com/sanalegon/twittor/models"
)

// funciones devuelven nada, metodos si. Esto si es una funcion
/* Registro es la funcion para crear en la BD el registro de usuario */
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) // decodificamos lo que viene en el body dentro del modelo t. El body es un stream, es un dato que solo se puede leer una vez. Una vez leido, una vez destruido

	if err != nil {
		http.Error(w, "Error en los datos recibidos => "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}

	// la funcion devuelve tres valores, solo que ignorare el primero y el ultimo
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese E-Mail", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar registrar el usuario => "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro insertar el registro el usuario => ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated) // Devolvemos por el header, que se logro crear el usuario. Se usara las constante de http para el codigo de estado
}
