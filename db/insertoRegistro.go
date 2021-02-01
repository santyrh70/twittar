package db

import (
	"context"
	"time"

	"github.com/santyrh70/twittar/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsetoResgistro */
func InsertoRegistro(u models.Ususario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("c-twittar")
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
