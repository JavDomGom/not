package db

import (
	"context"
	"time"

	"github.com/JavierDominguezGomez/not/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertUserRegister Function to record the registration of a new user in the database. */
func InsertUserRegister(u models.User) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("test_not")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
