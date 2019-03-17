package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*Tag - структура для тэгов, для группировки статей воедино*/
type Tag struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`              // идентификатор тэга в бд
	Name      string        `bson:"name" json:"tagname"`        // навзание тэга
	AuthorID  string        `bson:"id_author" json:"id_author"` // идентификатор автора тэга
	Author    *User         `bson:"author" json:"author"`       // автор тэга
	CreatedAt time.Time     `bson:"createdat" json:"createdat"` // дата создания тэга
}
