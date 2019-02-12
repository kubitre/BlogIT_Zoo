package Dao

import (
	"github.com/kubitre/blog/Models"
	"gopkg.in/mgo.v2/bson"
)

/*SettingTag - it is structure for setting bd connection
- структура для подключения к бд
*/
type SettingTag struct {
	Server   string
	Database string
}

const (
	collectionTags = "tags"
)

/*InsertDb - function for creating new line in dv and return error if this exist
- создание нового тега в бд и возврат ошибки, в случае возникновения*/
func (setting *SettingTag) InsertDb(Tag Models.Tag) (err error) {
	Tag.ID = bson.NewObjectId()
	err = db.C(collectionTags).Insert(&Tag)
	return
}

/*FindByID - function for finding element in db by id and returb this element and error if this exist
- поиск элемента в бд по его индентификатору и возврат его и ошибки, в случае её возникновения*/
func (setting *SettingTag) FindByID(id string) (tag Models.Tag, err error) {
	err = db.C(collectionTags).FindId(bson.ObjectIdHex(id)).One(&tag)
	return
}

/*FindAll - function for finding all elements in db collection and return slice with them and error if this exist
-  поиск всех записей коллекции в бд и возврат слайса с ними и ошибки, в слечае её возникновения*/
func (setting *SettingTag) FindAll() (tags []Models.Tag, err error) {
	err = db.C(collectionTags).Find(bson.M{}).All(&tags)
	return
}

/*Update - function for update line by indentificator for transmitted object and return error if this exist
- обновление записи в коллекции бд по её индетификатору новыми полями объекта и возрат ошибки в случае её возникновения*/
func (setting *SettingTag) Update(Tag Models.Tag) (err error) {
	err = db.C(collectionTags).UpdateId(Tag.ID, &Tag)
	return
}

/*Delete - function for delete line by indentificator in collection db and return error if this exist
- удаление элемента из бд по его индентификатору и возврат ошибки в случае её возникновения*/
func (setting *SettingTag) Delete(Tag Models.Tag) (err error) {
	err = db.C(collectionTags).Remove(&Tag)
	return
}
