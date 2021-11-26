package bd

import (
	"context"
	"time"

	"github.com/sanalegon/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Mandamos el id por separado, para no tener que procesar el id dentro de usuario
/* ModificoRegistro permite modificar el perfil del usuario */
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	// interfaz vacia
	registro := make(map[string]interface{}) // make permite crear slices o mapas. Aqui hacemos que todo el mapa sea de tipo interface

	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}

	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}

	registro["fechaNacimiento"] = u.FechaNacimiento // este dato lo forzamos sin verificacion

	if len(u.Avatar) > 0 {
		registro["Avatar"] = u.Avatar
	}

	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}

	if len(u.Biografia) > 0 {
		registro["bbiografia"] = u.Biografia
	}

	if len(u.Ubicacion) > 0 {
		registro["ubicacion"] = u.Ubicacion
	}

	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	// armamos el registro de actualizacion
	// Para actualizar un registro se usa el simbo peso y luego la palabra set
	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID) // convertir id string a object id

	// Usar filtro para solo actualizar a un  registro en especifico y no a todos
	// $eq => equal ==
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)

	if err != nil {
		return false, err
	}

	return true, nil

}
