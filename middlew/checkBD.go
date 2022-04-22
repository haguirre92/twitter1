package middlew

import (
	"net/http"

	"github.com/haguirre92/twitter1/bd"
)

func CheckBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
