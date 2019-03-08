package domain

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Entry domain model
type Entry struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Source    string        `bson:"source" json:"source"`
	CreatedAt time.Time     `bson:"created_at" json:"createdAt"`
}
