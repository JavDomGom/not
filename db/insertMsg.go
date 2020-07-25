package db

import (
	"context"
	"time"

	"github.com/JavierDominguezGomez/not/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertMsg records a msg in database. */
func InsertMsg(t models.RecordMsg) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("test_not")
	col := db.Collection("msg")

	register := bson.M{
		"userId":   t.UserID,
		"msg":      t.Msg,
		"datetime": t.Datetime,
	}
	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
