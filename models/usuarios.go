package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// El ID de mongo es un objeto binario. Es un slice de bits
// Para `, presionar alt + 96
// En base de datos, todo se guarda en formato bson
// _id convecion para los objectid de mongo
// omitempty, es para que lo omita si uno de los campos esta vacio y no sera tomado en cuenta para formar un json
// el formato json es lo que me devuelve. En vez de devolver _id, devuelve el id en minuscula
/* Usuario es el modelo de usuario de la base de MongoDB */
type Usuario struct {
	ID              primitive.ObjectID `bson: "_id, omitempty" json: "id"`
	Nombre          string             `bson: "nombre" json: "nombre, omitempty"` // No lo devolvera en json si no lo encuentra. Esto lo hacemos porque el nombre lo tendremos siempre
	Apellidos       string             `bson: "apellidos" json: "apellidos, omitempty"`
	FechaNacimiento time.Time          `bson: "fechaNacimiento" json: "fechaNacimiento, omitempty"`
	Email           string             `bson: "email json: "email"`                   // Aqui siempre devolvera el email
	Password        string             `bson: "password" json: "password, omitempty"` // Siempre omitempty porque nunca puedo regresar un pass por el navegador
	Avatar          string             `bson: "avatar" json: "avatar, omitempty"`
	Banner          string             `bson: "banner" json: "banner, omitempty"`
	Biografia       string             `bson: "biografia" json: "biografia, omitempty"`
	SitioWeb        string             `bson: "sitioWeb" json: "sitioWeb, omitempty"`
	Ubicacion       string             `bson: "ubicacion" json: "ubicacion, omitempty"`
}
