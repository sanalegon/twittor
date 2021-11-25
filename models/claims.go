package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Estructura en mayuscula porque sera exportada
/*Claim es la estructura usada para procesar el JWT*/
type Claim struct {
	Email              string             `json:"email"`
	ID                 primitive.ObjectID `bson:"_id" json:"id,omitempty"` // Viene con otro formato, y debemos trabajarlo con bson
	jwt.StandardClaims                    // Por ahora con estructura sin detallar, ya que nosotros no vamos trabajar con esos datos. Un ejemplo es la fecha de expiracion del token, pero eso no nos interesa.
}
