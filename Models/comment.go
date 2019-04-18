package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*Comment - структура данных для комментария в блоге*/
type Comment struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`                  // идентификатор комментария в бд
	CreatedAt   time.Time     `bson:"createdat" json:"createdat"`     // дата и время создания комментария
	Body        string        `bson:"commentbody" json:"body"`        // текст комментария
	Author      *User         `bson:"author" json:"author"`           // автор комментария
	AuthorID    string        `bson:"id_author" json:"id_author"`     // идентификатор автора комментария
	Verificated bool          `bson:"verificated" json:"verificated"` // разрешение на отображение комментарий(предварительно валидация комментария админом)
	ArticleID   string        `bson:"id_article" json:"id_article"`
}
