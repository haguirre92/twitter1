package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/haguirre92/twitter1/bd"
	"github.com/haguirre92/twitter1/models"
)

var Email string
var IDUsuario string

func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("laclavemas_Tesaquemehecreado")

	claims := &models.Claim{}
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}
	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		_, find, _ := bd.CheckExistUser(claims.Email)
		if find == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, find, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
