package Dao

import (
	"time"

	"github.com/kubitre/blog/Models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*TokenDao - структура для настройки коллекции tokens*/
type TokenDao struct {
	Database *mgo.Database
}

const (
	collectionTokens = "tokens"
)

/*CreateNewToken - запись нового токена в базу данных*/
func (setDao *TokenDao) CreateNewToken(newtoken string, id_user bson.ObjectId) (Models.Token, error) {

	tok := &Models.Token{
		ID:         bson.NewObjectId(),
		Value:      newtoken,
		ValidateTo: time.Second * 60,
		UserID:     id_user,
	}

	err := setDao.Database.C(collectionTokens).Insert(&tok)
	return *tok, err
}

/*FindToken - поиск токена в бд*/
func (setDao *TokenDao) FindToken(id bson.ObjectId) (*Models.Token, error) {
	var token Models.Token

	err := setDao.Database.C(collectionTokens).Find(bson.M{
		"id_user": id,
	}).One(&token)

	return &token, err
}

/*FindTokenByValue - поиск токена в бд по его значению*/
func (setDao *TokenDao) FindTokenByValue(tokenValue string) (*Models.Token, error) {
	var token Models.Token

	err := setDao.Database.C(collectionTokens).Find(bson.M{
		"maintoken": tokenValue,
	}).One(&token)

	return &token, err
}

/*RemoveToken - удаление токена из бд*/
func (setDao *TokenDao) RemoveToken(token *Models.Token) error {
	err := setDao.Database.C(collectionTokens).RemoveId(bson.M{
		"_id": token.ID,
	})

	return err
}
