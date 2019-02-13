package Routes

import (
	"github.com/gorilla/mux"
	"github.com/kubitre/blog/Config"
	"github.com/kubitre/blog/Dao"
	"github.com/kubitre/blog/Routes"
)

/*ApplicationForTesting - ir is structure for testing main api*/
type ApplicationForTesting struct {
	Database string
	Server   string
	Router   *mux.Router
}

/*Confugrator - it is function for configuration main router for testing module*/
func (sett *ApplicationForTesting) Confugrator() {
	var conf = Config.Configuration{}
	conf.Read()

	var daoServer = Dao.SettingDao{Server: conf.Server, Database: conf.Database}
	daoServer.Connect()

	sett.Router = Routes.StartSettingRoutes(conf, true)
	sett.Database = conf.Database
	sett.Server = conf.Server
}
