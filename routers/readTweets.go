package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/haguirre92/twitter1/bd"
)

func ReadTweets(writer http.ResponseWriter, reader *http.Request) {
	ID := reader.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writer, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	if len(reader.URL.Query().Get("page")) < 1 {
		http.Error(writer, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(reader.URL.Query().Get("page"))
	if err != nil {
		http.Error(writer, "Debe enviar el parametro pagina con un valor mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(page)
	response, right := bd.ReadTweets(ID, pag)
	if right == false {
		http.Error(writer, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(response)
}
