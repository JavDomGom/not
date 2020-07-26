package db

import (
	"context"
	"time"

	"github.com/JavierDominguezGomez/not/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*GetFollowersMsg Gets messages from my followers. */
func GetFollowersMsg(ID string, page int) ([]models.ReturnFollowersMsg, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("test_not")
	col := db.Collection("relation")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "msg",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "msg",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$msg"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"msg.datetime": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, conditions)
	var result []models.ReturnFollowersMsg
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
