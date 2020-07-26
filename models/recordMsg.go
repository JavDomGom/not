package models

import "time"

/*RecordMsg Structure for message. */
type RecordMsg struct {
	UserID   string    `bson:"userId" json:"userId,omitempty"`
	Msg      string    `bson:"msg" json:"msg,omitempty"`
	Datetime time.Time `bson:"datetime" json:"datetime,omitempty"`
}
