package Dao

import (
	"github.com/kubitre/blog/Models"
	"gopkg.in/mgo.v2/bson"
)

/*SettingComment - it is a structure for connection db for model Comment*/
type SettingComment struct {
	Server   string
	Database string
}

const (
	collectionComments = "comments"
)

/*InsertDb - function for creating new line in database, on input enter article object
- вставка нового элемента в бд*/
func (setting *SettingComment) InsertDb(Comment Models.Comment) (err error) {
	Comment.ID = bson.NewObjectId()
	err = db.C(collectionComments).Insert(&Comment)
	return
}

/*FindByID - function for finding liine in database, on input enter comment model with id which need find and return this object if exist in db or error object
- функция для поиска элемента в бд по индентификатору и возвратаа этого элемента или ошибки, в случае возникновения*/
func (setting *SettingComment) FindByID(id string) (comm Models.Comment, err error) {
	err = db.C(collectionComments).FindId(bson.ObjectIdHex(id)).One(comm)
	return
}

/*FindAll - function for find all lines in db and return slices with them with error or withour error
- функция для поиска всех записей в бд и возврат слайса с этими элементами а также ошибки, если таковая возникает*/
func (setting *SettingComment) FindAll() (comms []Models.Comment, err error) {
	err = db.C(collectionComments).Find(bson.M{}).All(comms)
	return
}

/*Update - function for update line in db by id and return error without error
- функция для обнлвления записи в бд по индентификатору и возвращает ошибку при возникновении*/
func (setting *SettingComment) Update(Comment Models.Comment) (err error) {
	err = db.C(collectionComments).UpdateId(Comment.ID, &Comment)
	return
}

/*Delete - remove line in db by id and return with error or without
- Удаление комментария по индентификатору и возврат ошибки в случае возникновении*/
func (setting *SettingComment) Delete(Comment Models.Comment) (err error) {
	err = db.C(collectionComments).Remove(&Comment)
	return
}
