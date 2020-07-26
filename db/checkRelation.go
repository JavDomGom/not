package db

import (
	"context"
	"fmt"
	"time"

	"github.com/JavierDominguezGomez/not/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckRelation Check for relation beetween two users in database. */
func CheckRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("test_not")
	col := db.Collection("relation")

	condition := bson.M{
		"userid":         t.UserID,
		"userrelationid": t.UserRelationID,
	}

	var result models.Relation
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
