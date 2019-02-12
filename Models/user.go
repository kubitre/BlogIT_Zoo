package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*User - it is structure for main type user in our system*/
type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Username    string        `bson:"name" json:"username"`
	Password    string        `bson:"password" json:"password"`
	Email       string        `bson:"email" json:"email"`
	Verificated bool          `bson:"verificated" json:"verificated"`
	CreatedAt   time.Time     `bson:"createdat" json:"createdat"`
}
