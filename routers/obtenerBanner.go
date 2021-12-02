package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/sanalegon/twittor/bd"
)

/* ObtenerBanner envia el banner al HTTP */
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + perfil.Banner)

	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	// Envio de la imagen al http y la guardamos en w. Esto no tiene prueba en postman, porque el frontend mostrara la imagen directamente y no me devolvera ningun dato para el backend
	_, err = io.Copy(w, OpenFile) // enviamos el archivo en modo binario al response writer

	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
	// no retornamos estado 201 porque el frontend solo necesita saber si le llego la imagen o no
}
