package Dao

import (
	"github.com/kubitre/blog/Models"
	"gopkg.in/mgo.v2/bson"
)

/*SettingArticle it is a structure for setting connection to db
- структура для настройки подключения к бд
*/
type SettingArticle struct {
	Server   string
	Database string
}

const (
	collectionArticles = "articles"
)

/*InsertDb - function for creating new line in database, on input enter article object
- вставка нового элемента в бд*/
func (setting *SettingArticle) InsertDb(article Models.Article) (art Models.Article, err error) {
	art.ID = bson.NewObjectId()
	err = db.C(collectionArticles).Insert(&article)
	return
}

/*FindByID - function for finding line in databsae, in input enter article object with id on bson.ObjectId type and return his if founded and nil else
- поиск элемента по его индентификатору и возврат его либо ошибки, в случае возникновении*/
func (setting *SettingArticle) FindByID(id string) (art Models.Article, err error) {
	err = db.C(collectionArticles).FindId(bson.ObjectIdHex(id)).One(&art)
	return
}

/*FindAll - function for finding all lines in database and return slice with them or error else
- поиск всех записей в коллекции и возврат слайса с ними лио ошибки, в случае возникновения*/
func (setting *SettingArticle) FindAll() (arts []Models.Article, err error) {
	err = db.C(collectionArticles).Find(bson.M{}).All(&arts)
	return
}

/*Update - function for update line in database and return nil error or with Error
- обновление записей по индентификатору с заданным полями и возврат ошибки, в случае её возникновения*/
func (setting *SettingArticle) Update(article Models.Article) (err error) {
	err = db.C(collectionArticles).UpdateId(article.ID, &article)
	return
}

/*Delete - function for delete line in database and return nil error or with Error
- удаления элемента по его индентификатору в бд и возрат ошибки, в случае её возникновения*/
func (setting *SettingArticle) Delete(article Models.Article) (err error) {
	err = db.C(collectionArticles).Remove(&article)
	return
}
