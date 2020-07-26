package db

import (
	"context"
	"fmt"
	"time"

	"github.com/JavierDominguezGomez/not/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*GetAllUsers Get all users who have a relation with me. */
func GetAllUsers(ID string, page int64, search string, userType string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("test_not")
	col := db.Collection("users")

	var results []*models.User

	opts := options.Find()
	opts.SetSkip((page - 1) * 20)
	opts.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, opts)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, include bool

	for cursor.Next(ctx) {
		var s models.User
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		include = false

		found, err = CheckRelation(r)
		if userType == "new" && !found {
			include = true
		}
		if userType == "follow" && found {
			include = true
		}
		if r.UserRelationID == ID {
			include = false
		}

		if include {
			s.Password = ""
			s.Biography = ""
			s.WebSite = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	cursor.Close(ctx)

	return results, true
}
