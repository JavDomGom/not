package db

import (
	"context"
	"time"

	"github.com/JavierDominguezGomez/not/models"
)

/*InsertFollow record relation in database. */
func InsertFollow(t models.Follow) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("not")
	col := db.Collection("follows")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
