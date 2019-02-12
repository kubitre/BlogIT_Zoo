package Models

import (
	"gopkg.in/mgo.v2/bson"
)

/*Article type for all layers*/
type Article struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Tags        []*Tag        `bson:"tags" json:"tags"`
	Autor       *User         `bson:"author" json:"author"`
}
