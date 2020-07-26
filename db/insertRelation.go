package db

import (
	"context"
	"time"

	"github.com/JavierDominguezGomez/not/models"
)

/*InsertRealtion record relation in database. */
func InsertRealtion(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("test_not")
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
