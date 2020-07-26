package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ReturnFollowersMsg Structure for returned message. */
type ReturnFollowersMsg struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userId"`
	UserRelationID string             `bson:"userrelationid" json:"userRelationId"`
	Msg            struct {
		Msg      string    `bson:"msg" json:"msg"`
		Datetime time.Time `bson:"datetime" json:"datetime,omitempty"`
		ID       string    `bson:"_id" json:"_id,omitempty"`
	}
}
