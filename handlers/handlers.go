package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/haguirre92/twitter1/middlew"
	"github.com/haguirre92/twitter1/routers"
	"github.com/rs/cors"
)

func Managers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.CheckBD(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8083"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
