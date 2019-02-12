package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Tag struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"tagname"`
	Author    *User         `bson:"author" json:"author"`
	CreatedAt time.Time     `bson:"createdat" json:"createdat"`
}
