package bd

import (
	"context"
	"time"

	"github.com/sanalegon/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertoRegistro es la parada final con la BD para insertar los datos del usuario */
func InsertoRegistro(u models.Usuario) (string, bool, error) {

	// el contexto background es lo que vengo trabajando en la BD. Es el famoso TODO. Le hare una actualizacion, agregandole un contexto en miniatura para el timeout
	// Primero digo en que contexto grabare y luego cuanto el el tiempo para el timeout (15 * segundos es de 15 segundos multiplicado 1 time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// defer es una instruccion que se setea al comienzo, pero que se ejecuta justo al final de la funcion
	defer cancel() // Al objeto devuelto por ctx, le hago un cancel. Cancela el timeout/contexto

	db := MongoCN.Database("twittor")
	// coleccion
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password) // esta funcion estara en otra parte, porque es muy reutilizable. Mejor tenerlo facil de agarrar

	result, err := col.InsertOne(ctx, u) // a la coleccion insertaremos un registro, con el contexto de 15 segundos (para que no tarde mas que eso) y mandamos el modelo u

	if err != nil {
		return "", false, err // "" es devolver un ID vacio, false es que la funcion no fue satisfactoria
	}

	// No lo usaremos, pero es util para saber como objener la ID del usuairo creado
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	// Es necesario convertir a string el ID (era boleano)
	return ObjID.String(), true, nil
}
