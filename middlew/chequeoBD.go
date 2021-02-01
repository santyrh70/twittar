package middlew

import (
	"net/http"

	"github.com/santyrh70/twittar/db"
)

/*ChequeoDB es el middlew 	ue me permite conocer estado de db*/
func ChequeoDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexion perdida con la db", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
