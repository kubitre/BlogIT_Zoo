package Dao

import (
	"github.com/kubitre/blog/Models"
	"gopkg.in/mgo.v2/bson"
)

/*SettingUser - it is structure for connection to db*/
type SettingUser struct {
	Server   string
	Database string
}

const (
	collectionUsers = "users"
)

/*InsertDb - function for creating new user in db*/
func (setting *SettingUser) InsertDb(User Models.User) (err error) {
	User.ID = bson.NewObjectId()
	err = db.C(collectionUsers).Insert(&User)
	return
}

/*FindByID - function for finding element by indentificator and return this object with error if this exist*/
func (setting *SettingUser) FindByID(id string) (usr Models.User, err error) {
	err = db.C(collectionUsers).FindId(bson.ObjectIdHex(id)).One(usr)
	return
}

/*FindAll - function for finding all elements in collection db and return slice with them and error if this exist*/
func (setting *SettingUser) FindAll() (usrs []Models.User, err error) {
	err = db.C(collectionUsers).Find(bson.M{}).All(&usrs)
	return
}

/*Update - function for update element in db collection by indentificator and return error if this exist*/
func (setting *SettingUser) Update(User Models.User) (err error) {
	err = db.C(collectionUsers).UpdateId(User.ID, &User)
	return
}

/*Delete - function for delete element by indentificator and return error if this exist*/
func (setting *SettingUser) Delete(User Models.User) (err error) {
	err = db.C(collectionUsers).Remove(&User)
	return
}
