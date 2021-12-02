package bd

import (
	"context"
	"time"

	"github.com/sanalegon/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* LeoTweetsSeguidores lee los tweets de mis seguidores */
func LeoTweetsSegudores(ID string, pagina int) ([]models.DevuelvoTweetsSeguidores, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20

	condiciones := make([]bson.M, 0) // hacemos un slice de timo bson con 0 elementos

	// param uno lo que adicionamos, param dos son las condiciones. Match es un comando para buscar el usuario id de mi relacion. Luego decimos que usuario id debe ser igual al ID (que viene por parametro)
	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	// expresion lookup es para unir dos tablas. Cogeremos lo que filtramos en la linea arriba y lo unimos a la tabla tweets
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",             // tabla con la que queremos unir la relacion
			"localField":   "usuariorelacionid", // campo por el que vamos a unir
			"foreignField": "userid",            // como se llama el localfield en la tabla from. en este caso tweet
			"as":           "tweet",             // alias para la tabla
		}})

	// unwind permite procesar los resultados. Nos permite que todos los documentos nos vengan exactamente iguales
	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	// 1 es organizada de mayor a menor. -1 es el contrario
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)
	// no usaremos el cursor ya que aggregate nos ayuda a no tener que procesarlo el cursor. no me toca recorrer uno por uno el cursor
	var result []models.DevuelvoTweetsSeguidores

	// chequemos si hay error en el cursor. Procesamos todos lo registros de una y si hay error, error recibira un valor
	err = cursor.All(ctx, &result) // mandamos la conexion y luego el resultado lo mandara  a la variable result formateando con las reglas de devuelvotweetsseguidores
	if err != nil {
		return result, false
	}

	return result, true
}
