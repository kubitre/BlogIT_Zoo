package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*Article type for all layers*/
type Article struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`              // идентификатор статьи в бд
	Name        string        `bson:"name" json:"header"`         // навзание статьи
	Description string        `bson:"description" json:"preview"` // описание статьи (неполноая статья)
	Body        string        `bson:"body" json:"content"`        // основное содержание статьи
	TagsIDs     []string      `bson:"-" json:"id_tags"`           // идентификаторы тегов
	Tags        []Tag         `bson:"tags" json:"tags"`           // теги
	Autor       *User         `bson:"author" json:"author"`       // автор статьи
	AuthorID    string        `bson:"-" json:"id_author"`         // идентификатор автора статьи
	Likes       []string      `bson:"-" json:"id_likes"`          // идентификаторы лайков
	
	CreatedAt   time.Time     `bson:"createdat" json:"createdat"` // дата и время создания статьи
}
