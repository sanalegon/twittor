package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/sanalegon/twittor/bd"
	"github.com/sanalegon/twittor/models"
)

// Tiene tratamiento parecido a subir un banner, pero se mantiene en otro archivo por orden
/* SubirAvatar sube el avatar al servidor */
func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	// capturar del request el archivo
	// con formfile decimos que lo procesaremos como un archivo de html y dentro de este habra el archivo avatar
	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]               // capturamos solo la extension del archivo. Ponemos [1] para capturar el elemento 1 y de una coger la extension
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension // seteamos donde estara el archivo

	// Abrimos el archivo. Param 1: archivo, Param 2: constante que en este caso dice que es solo para escribir y despues del | es otra costante, param 3: permisos 0666 => lectura, escritura y ejecucion permisos
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	// f es donde voy a copiar y param 2 es lo que voy a copiar
	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	// modificamos el registro en la bd con la nueva locacion y nombre de archivo
	var usuario models.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + extension

	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la BD ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
