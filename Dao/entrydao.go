package Dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

/*SettingDao - it is structure for setting data access objects for bd connection
- структура для настройки подключения к бд*/
type SettingDao struct {
	Server   string
	Database string
	Db       *mgo.Database
}

/*Connect - function for connection to database server*/
func (m *SettingDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	m.Db = session.DB(m.Database)
}
