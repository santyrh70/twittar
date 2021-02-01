package routers

import (
	"encoding/json"
	"net/http"

	"github.com/santyrh70/twittar/db"
	"github.com/santyrh70/twittar/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Ususario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email de usuario es requerido", 400)
		return
	}

	if len(t.Email) < 6 {
		http.Error(w, "La contraseña debe ser mayor a 6 caracteres", 400)
		return
	}

	_, encontrado, _ := db.ChequeoExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya esxiste un usuario registrado con este email", 400)
		return
	}

	_, status, err := db.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro insertar el usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
