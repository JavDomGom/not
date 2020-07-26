package db

import (
	"context"
	"time"

	"github.com/JavierDominguezGomez/not/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ModifyRegister modify the user's profile. */
func ModifyRegister(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("test_not")
	col := db.Collection("users")

	register := make(map[string]interface{})
	if len(u.Name) > 0 {
		register["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		register["lastName"] = u.LastName
	}
	if !u.DateOfBirth.IsZero() {
		register["dateOfBirth"] = u.DateOfBirth
	}
	if len(u.Avatar) > 0 {
		register["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		register["banner"] = u.Banner
	}
	if len(u.Biography) > 0 {
		register["biography"] = u.Biography
	}
	if len(u.Location) > 0 {
		register["location"] = u.Location
	}
	if len(u.WebSite) > 0 {
		register["webSite"] = u.WebSite
	}

	updtString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}
