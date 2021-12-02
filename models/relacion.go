package models

/* Relacion modelo para grabar la relacion de un usuario con otro*/
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`                 // Mi id
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"` // Id del usuario que estoy siguiendo
}
