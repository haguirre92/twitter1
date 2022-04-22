package main

import (
	"log"

	"github.com/haguirre92/twitter1/bd"
	"github.com/haguirre92/twitter1/handlers"
)

func main() {
	if bd.CheckConection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Managers()
}
