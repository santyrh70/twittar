package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/santyrh70/twittar/db"
	"github.com/santyrh70/twittar/jwt"
	"github.com/santyrh70/twittar/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Ususario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Ususario o contrseña invalido"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del ususario es requerido", 400)
		return
	}
	documento, existe := db.IntentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario o contraseña invalidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurio un error al intentar general el token"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	/*como grabar cookie*/
	expirationTime := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
