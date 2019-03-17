package Dao

import (
	"github.com/kubitre/blog/Models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*TokenDao - структура для настройки коллекции tokens*/
type TokenDao struct {
	Database *mgo.Database
}

const (
	collectionTokens = "/tokens"
)

/*CreateNewToken - запись нового токена в базу данных*/
func (setDao *TokenDao) CreateNewToken(newtoken Models.Token) (Models.Token, error) {
	newtoken.ID = bson.NewObjectId()
	err := setDao.Database.C(collectionTokens).Insert(&newtoken)
	return newtoken, err
}

/*FindToken - поиск токена в бд*/
func (setDao *TokenDao) FindToken(id string) error {
	return nil
}
