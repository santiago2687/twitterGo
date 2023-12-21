package bd

import (
	"context"

	"github.com/santiago2687/twitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx := context.TODO()

	db := MongoCN.Database(DatabaseName)
	coleccion := db.Collection("usuarios")

	condition := bson.M{"email": email}

	var resultado models.Usuario

	err := coleccion.FindOne(ctx, condition).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
