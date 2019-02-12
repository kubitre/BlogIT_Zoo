package Routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kubitre/blog/Config"
)

/*RouteSetting - Main route layer for settings all routes*/
type RouteSetting struct {
	apiVersion string
}

const (
	apiRouteMain = "/v"
	version      = "1"
)

/*GetVersion - function for getting version of api*/
func (routesetting *RouteSetting) GetVersion(w http.ResponseWriter, r *http.Request) {

}

/*GetStatus - function for getting status on backend*/
func (routesetting *RouteSetting) GetStatus(w http.ResponseWriter, r *http.Request) {

}

/*StartSettingRoutes - function for settings database and routes*/
func StartSettingRoutes(config Config.Configuration) {
	routerArticles := mux.NewRouter()

	routerArticles.HandleFunc(apiRouteMain + version + ArticleRoute.)
}
