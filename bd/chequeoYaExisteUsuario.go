package bd

import (
	"context"
	"time"

	"github.com/sanalegon/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Retornamos models, para que regrese todo el registro. Asi ya tengo estos datos por si cualquier cosa
/* ChequeoYaExisteUsuario recibe un emailde parametro y chequea si ya esta en la BD */
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor") // Decimos que bd usaremos.Devolvemos twittor
	col := db.Collection("usuarios")  // que coleccion usaremos

	// Es tipo bson. M es un formato tipo mapstring (formato json)
	condicion := bson.M{"email": email}

	var resultado models.Usuario // tipo usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado) //lo que devuelva findone, lo convertimos en json(decode) dentro de resultado, siende este ultimo un puntero
	ID := resultado.ID.Hex()                              // resultado tendria el ID como tipo objectid, para no tener que trabajarlo con ese tipo, lo convertimos en hexadecimal en formato string

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
