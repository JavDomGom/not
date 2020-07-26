package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ReturnMsg Structure to return each message. */
type ReturnMsg struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID   string             `bson:"userid" json:"userId,omitempty"`
	Msg      string             `bson:"msg" json:"msg,omitempty"`
	Datetime time.Time          `bson:"datetime" json:"datetime,omitempty"`
}
