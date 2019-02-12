package main

import (
	"fmt"

	"github.com/kubitre/blog/Dao"

	"github.com/kubitre/blog/Config"
	"github.com/kubitre/blog/Routes"
)

func init() {
	var conf = Config.Configuration{}
	conf.Read()

	var daoServer = Dao.SettingDao{Server: conf.Server, Database: conf.Database}
	daoServer.Connect()

	Routes.StartSettingRoutes(conf)
	fmt.Println("api was started")
}

func main() {
	fmt.Println("Api was started")
}
