package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/sanalegon/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* LeoUsuario lee los usuarios registrados en el sistema, si se recibe "R" en quienes trae solo los que se relacionan conmigo*/
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()

	// Seteamos propiedades a findOptions. El orden importa, setskip no puede ir luego de setlimit
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search}, // usamos regex debido a que viene un string a buscar. Regex es una expresion regular. La i indica que no se fija si es mayus o minus
	}

	// cur(sor). Cuando no uso findone, el resultado me lo devuelve en un cursor
	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error() + " => error while finding results")
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {
		fmt.Println("exploring cur ...")
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error() + " => error while exploring cur in models usuarios")
			return results, false
		}

		// r(elacion)
		var r models.Relacion
		r.UsuarioID = ID // No usare la global para respetar lo que nos llega como parametro
		fmt.Println("relacion id: " + s.ID.Hex())
		r.UsuarioRelacionID = s.ID.Hex() // Extraemos la parte string del usuario

		incluir = false // Para saber si debo incluir el usuario de  la actual iteraccion en respuesta o no

		encontrado, err = ConsultoRelacion(r) // no necesito bd. ya que ahora estoy parado en la carpeta bd
		// new seria para todos los usuarios que no estoy siguiendo
		if tipo == "new" && encontrado == false {
			incluir = true
		}
		// Aqui solo quiero el listado de la gente que sigo
		if tipo == "follow" && encontrado == true {
			incluir = true
		}
		fmt.Println("id 1: "+r.UsuarioRelacionID+", id 2:"+ID, incluir)
		// Por si se da el fe de erratas de mandar el mismo usuario que los que se desea consultar
		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		fmt.Println("new user: " + s.Nombre)
		fmt.Println("And will be appended ? %t", incluir)

		if incluir == true {
			// Blanqueo todos los atributos que no necesito consultar. Sirve de seguridad y no mandar info inutil
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""

			// adicionamos donde se graba (param 1) y param 2 es lo que grabamos. Que es un puntero de memoria. Ira a el y extraera la info del puntero de memoria
			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error() + " => error in cur")
		return results, false
	}

	// cerrar cursor es muy importante
	cur.Close(ctx)
	return results, true
}
