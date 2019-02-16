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

/*ConfigureDbConnection - confgurating db connection*/
func (sett *ApplicationForTesting) ConfigureDbConnection() {
	var conf = Config.Configuration{}
	conf.Read()

	var daoServer = Dao.SettingDao{Server: conf.Server, Database: conf.Database}
	daoServer.Connect()

	sett.Database = conf.Database
	sett.Server = conf.Server
}

/*Configurator - it is function for configuration main router for testing module*/
func (sett *ApplicationForTesting) Configurator(typeroute string) (*mux.Router, error) {

	sett.Router = mux.NewRouter()

	switch typeroute {
	case "articles":
		sett.ConfigureArticleRoute()
	case "users":
		sett.ConfigureUserRoute()
	case "tags":
		sett.ConfigureTagRoute()
	case "comments":
		sett.ConfigureCommentRoute()
	default:
		return nil, nil
	}

	return sett.Router, nil
}

/*ConfigureArticleRoute - it is function configureting mux router by routes with articles for testing system*/
func (sett *ApplicationForTesting) ConfigureArticleRoute() {
	Routes.StartSettingRouterArticle(sett.Router)
}

/*ConfigureCommentRoute - it is function configurating mux router by routes with comments for testing system*/
func (sett *ApplicationForTesting) ConfigureCommentRoute() {
	Routes.StartSettingRouterComment(sett.Router)
}

/*ConfigureTagRoute - it is function configurating mux router by routes with tags for testing system*/
func (sett *ApplicationForTesting) ConfigureTagRoute() {
	Routes.StartSettingRouterTag(sett.Router)
}

/*ConfigureUserRoute - it is function configurating mux router by routes wuth users for tsting system*/
func (sett *ApplicationForTesting) ConfigureUserRoute() {
	Routes.StartSettingRouterUser(sett.Router)
}
