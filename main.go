package main

// Para traer el handlers, hay que poner primero nuestro user en github. Esto no significa que lo va a buscar en github, solo que asi se importa archivos locales en golang
import (
	"log"

	"github.com/sanalegon/twittor/bd"
	"github.com/sanalegon/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
