package bd

import (
	"context"

	"github.com/santiago2687/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(usuario models.Usuario) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	coleccion := db.Collection("usuarios")

	usuario.Password, _ = EncriptarPassword(usuario.Password)

	result, err := coleccion.InsertOne(ctx, coleccion)

	if err != nil {
		return "", false, err
	}

	ObjectId, _ := result.InsertedID.(primitive.ObjectID)

	return ObjectId.String(), true, nil
}
