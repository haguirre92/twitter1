package routers

import (
	"encoding/json"
	"net/http"

	"github.com/haguirre92/twitter1/bd"
)

func WatchPerfil(writer http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.SearchPerfil(ID)
	if err != nil {
		http.Error(writer, "Ocurrio un error al buscar el registro "+err.Error(), 400)
		return
	}

	writer.Header().Set("context-type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(perfil)
}
