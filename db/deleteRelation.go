package db

import (
	"context"
	"time"

	"github.com/JavierDominguezGomez/not/models"
)

/*DeleteRelation Delete a specific relation. */
func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("test_not")
	col := db.Collection("relation")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
