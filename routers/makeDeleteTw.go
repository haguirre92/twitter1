package routers

import (
	"net/http"

	"github.com/haguirre92/twitter1/bd"
)

func MakeDeleteTw(writer http.ResponseWriter, reader *http.Request) {
	ID := reader.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(ID, IDUsuario)
	if err != nil {
		http.Error(writer, "Ocurrio un error borrando el tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-type", "application-json")
	writer.WriteHeader(http.StatusCreated)
}
