package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:toor@c-twittar.qp41z.mongodb.net/<dbname>?retryWrites=true&w=majority")

/*ConectarDB me permite conectar a la base de datos*/
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion exitosa con la db")
	return client
}

/*CheckConnection verifica la conexion db*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
