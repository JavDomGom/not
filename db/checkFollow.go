package db

import (
	"context"
	"time"

	"github.com/JavierDominguezGomez/not/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckFollow Check for relation beetween two users in database. */
func CheckFollow(t models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("not")
	col := db.Collection("follows")

	condition := bson.M{
		"userId":         t.UserID,
		"userFollowedID": t.UserFollowedID,
	}

	var result models.Follow
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}
	return true, nil
}
