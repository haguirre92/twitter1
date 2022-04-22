package routers

import (
	"encoding/json"
	"net/http"

	"github.com/haguirre92/twitter1/bd"
	"github.com/haguirre92/twitter1/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es un campo obligatorio", 400)
		return
	}

	if len(t.Password) == 0 {
		http.Error(w, "La contraseña debe contener minimo 6 caracteres", 400)
		return
	}

	_, find, _ := bd.CheckExistUser(t.Email)

	if find == true {
		http.Error(w, "Ya existe un usuario con este email", 400)
		return
	}

	_, status, err := bd.InsertRegister(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro de usuario"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
