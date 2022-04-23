package routers

import (
	"enconding/json"
	"net/http"
	"time"

	"github.com/haguirre92/twitter1/bd"
	"github.com/haguirre92/twitter1/jwt"
	"github.com/haguirre92/twitter1/models"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "Usuario y/o contraseña invalidos"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(writer, "El email del usuario es requerido", 400)
		return
	}

	document, exist := bd.TryLogin(user.Email, user.Password)
	if exist == false {
		http.Error(writer, "Usuario y/o contraseña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GeratedJWT(document)
	if err != nil {
		http.Error(writer, "Ocurrio un error al intentar general el token correspondiente "+err.Error(), 400)
		return
	}

	resp := models.ResponseLogin{
		Token: jwtKey,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(writer, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
