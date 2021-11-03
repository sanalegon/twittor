package bd

import "golang.org/x/crypto/bcrypt"

/* EncriptarPassword es la rutinar que me permite encriptar la password */
func EncriptarPassword(pass string) (string, error) {
	costo := 8 // Es la cantidad de veces a encriptar. Formula 2 a la costo. Lo aconsejable es que minimo sea 6. 6 para usuarios normales, 8 para super users
	//mandamos un slice (vector sin cantidad de elementos). Es dinamico su tamaño. [] => slice, byte => tipo, pass => es el parametro
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)

	return string(bytes), err
}
