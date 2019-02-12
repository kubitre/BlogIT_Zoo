package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Comment struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	CreatedAt   time.Time     `bson:"createdat" json:"createdat"`
	Body        string        `bson:"commentbody" json:"body"`
	Author      *User         `bson:"author" json:"author"`
	Verificated bool          `bson:"verificated" json:"verificated"`
}
