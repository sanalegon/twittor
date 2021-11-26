package bd

import (
	"context"
	"log"
	"time"

	"github.com/sanalegon/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Devolvemos un slice tipo array con todos los tweets de una persona. Slice es para hacer dinamico la cantidad de elemento que va a tener. Es un slice de tipo models devuelvo tweets
/* LeoTweets lee los tweets de un perfil */
func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	// slice
	var resultados []*models.DevuelvoTweets

	condicion := bson.M{
		"userid": ID,
	}

	// options es un paquete para definir opciones y asi filtrar y determinar un comportamiento a mi consulta de base de datos
	opciones := options.Find() // decimos que trabajaremos con options en modo find

	// Ahora todo lo que setee, seran propiedades que intervendran durante el find
	opciones.SetLimit(20)                               // cuantos me debe traer limite por pagina (tema de paginacion)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}}) //Decimos como vendran organizados los tweets. Bson.D es otro formato de bson. Aqui traera todos los documentos ordenados por la fecha(key) en -1, que siginifica en orden descendente
	opciones.SetSkip((pagina - 1) * 20)                 // lo que hago aqui es que en pagina 1, no salto nada. En pagina 2 salto 20. En pagina 3 salto 40, y asi sucesivamente. Cada pagina contiene 20 registros. Ejm: para pagina 1: (1 - 1) * 20 = 0 * 20 = 0 => salta cero

	// cursor es un puntero tipo table de base de datos en donde se guardaran los resultados y podre recorrerlos uno a la vez
	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	// no usamos ctx para no mezclarlo con la bd. TODO es crear un contexto vacio
	for cursor.Next(context.TODO()) {
		// por cada iteracion crearemos un registro nuevo
		var registro models.DevuelvoTweets // aqui trabajo con cada tweet en particular
		// cursor es como un json y el resultado del decode se guardara en registro
		err := cursor.Decode(&registro)

		if err != nil {
			return resultados, false
		}
		// param 1: donde lo guardare
		// param 2: puntero a registro
		resultados = append(resultados, &registro) // sirve para agregar en un slice un elemento. Tambien podria usar make o asignacion directa
	}

	return resultados, true
}
