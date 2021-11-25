package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/sanalegon/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* BuscoPerfil busca un perfil en la BD */
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var perfil models.Usuario

	// convertir ID en objectID, porque asi lo tenngo en mi modelo
	objID, _ := primitive.ObjectIDFromHex(ID)

	// la condicionn a buscar sera igual a lo que figura dentro de la base de datos
	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)

	// limpiamos el valor de password para no devolverla. Esto se hace por temas de seguridad
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no enncontrado " + err.Error())
		return perfil, err
	}

	return perfil, nil
}
