package Dao

import (
	mgo "gopkg.in/mgo.v2"
)

/*SettingDao - it is structure for setting data access objects for bd connection
- структура для настройки подключения к бд*/
type SettingDao struct {
	Server   string
	Database string
}

var db *mgo.Database
