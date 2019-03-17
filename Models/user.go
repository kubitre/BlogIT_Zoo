package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*User - it is structure for main type user in our system*/
type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`              // идентификатор пользователя в бд
	Avatar      []byte        `bson:"avatar" json:"avatar"`       // аватар пользователя
	Username    string        `bson:"name" json:"username"`       // имя пользователя
	Password    string        `bson:"password" json:"-"`          // пароль пользователя
	Email       string        `bson:"email" json:"email"`         // email пользователя
	Verificated bool          `bson:"verificated" json:"-"`       // подвтерждение аккаунта
	CreatedAt   time.Time     `bson:"createdat" json:"createdat"` // дата и время создания аккаунта
}
