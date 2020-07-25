package models

import "time"

/*RecordMsg Structure of the message. */
type RecordMsg struct {
	UserID   string    `bson:"userId" json:"userId,omitempty"`
	Msg      string    `bson:"msg" json:"msg,omitempty"`
	Datetime time.Time `bson:"datetime" json:"datetime,omitempty"`
}
