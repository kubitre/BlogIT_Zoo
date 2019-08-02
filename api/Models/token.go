package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*Token - it is structure for access layer of our system*/
type Token struct {
	ID         bson.ObjectId `bson:"_id" json:"-"`               // идентификатор токена в бд
	Value      string        `bson:"maintoken" json:"token"`     // значение токена
	Createdat  time.Time     `bson:"createdat" json:"createdat"` // дата создания
	ValidateTo time.Duration `bson:"valideto" json:"-"`          // время валидности токена с момента создания
	UserID     bson.ObjectId `bson:"id_user" json:"-"`           // кому принадлежит токен
}