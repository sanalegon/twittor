package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// para guardar texto dentro del log de ejecucion

// MongoC es lo mas cercano a "variable global" que tenemos y sera usada en todos los archivos de la carpeta bd
/* MongoCN es el objeto de conexoin a la BD */
var MongoCN = ConectarBD()                                                                                                                   // variable que ejecuta la funcion
var clientOptions = options.Client().ApplyURI("mongodb+srv://sanalegon:1234@cluster0.tqs48.mongodb.net/twittor?retryWrites=true&w=majority") // var que me permite setear la url de la base de datos

/* Conectar es la funcion que me permite conectar la BD */
func ConectarBD() *mongo.Client {
	// Hacemos una conexion a la bd(connect), luego toma la conexion de clientOptions y tiene un contexto(context.TODO)
	// En go no existen variables globales, archivos de funciones, etc (no estan bien vistios)
	// Dentro del llamado de una api, go tiene algo llamado contexto. Esto es un espacio de memoria donde yo compartire cosa. Tipo setear un contexto de ejecucion (no mas de 15 segundos)
	// Contextos sirven para comunicar informacion entre ejecucion y ejecucion
	// TODO es lo basico del context. En context puedo crear tambien variables, porque estos son un entorno de ejecucion en donde puedo guardar info que estara disponible mientras que el cotexto este ejecutandose
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("error was found")
		log.Fatal(err.Error())
		return client
	}
	//arriba usa := y no abajo, porque err ya ha sido inicializada. := es solo cuando se inicializa por primera vez
	err = client.Ping(context.TODO(), nil) // ping para comprobar que la bd esta arriba

	if err != nil {
		log.Println("error when trying to ping db")
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion Exitosa con la BD")
	return client
}

/* ChequeoConnection es el Ping a la BD */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil) // ping para comprobar que la bd esta arriba

	if err != nil {
		log.Fatal(err)
		return 0
	}

	return 1
}
