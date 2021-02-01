package db

import (
	"context"
	"time"

	"github.com/santyrh70/twittar/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoExisteUsuario(email string) (models.Ususario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("c-twittar")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}
	var resultado models.Ususario
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
