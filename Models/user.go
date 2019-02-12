package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Username    string        `bson:"name" json:"username"`
	Password    string        `bson:"password" json:""`
	Email       string        `bson:"email" json:"email"`
	Verificated bool          `bson:"verificated" json:"verificated"`
	CreatedAt   time.Time     `bson:"createdat" json:"createdat"`
}
