package routers

import (
	"encoding/json"
	"net/http"

	"github.com/haguirre92/twitter1/bd"
	"github.com/haguirre92/twitter1/models"
)

func ModifyPerfil(writer http.ResponseWriter, request *http.Request) {
	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)

	if err != nil {
		http.Error(writer, "Datos incorrectos: "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModifyRegister(user, IDUsuario)

	if err != nil {
		http.Error(writer, "Ocurrio un error al modificar el registro. Intente de nuevo"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(writer, "No se ha logrado modificar el registro: "+err.Error(), 400)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
