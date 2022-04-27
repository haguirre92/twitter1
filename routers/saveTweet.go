package routers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/haguirre92/twitter1/bd"
	"github.com/haguirre92/twitter1/models"
)

func SaveTweet(writer http.ResponseWriter, request *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(request.Body).Decode(&message)
	log.Println("Estamos aca: ", IDUsuario)
	register := models.SaveTweet{
		UserID:      IDUsuario,
		Message:     message.Message,
		DateCreated: time.Now(),
	}

	_, status, err := bd.InsertTweet(register)
	if err != nil {
		http.Error(writer, "Error en la insercion del tweet"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(writer, "No se puso insertar el tweet", 400)
		return
	}

	writer.WriteHeader(http.StatusCreated)

}
