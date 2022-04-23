package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/haguirre92/twitter1/models"
)

func GeneratedJWT(user models.User) (string, error) {

	myKey := []byte("laclavemas_Tesaquemehecreado")
	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastName":  user.LastName,
		"dateBirth": user.DateBirth,
		"biografia": user.Biografia,
		"ubicacion": user.Ubicacion,
		"siteWeb":   user.SiteWeb,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil

}
