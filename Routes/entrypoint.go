package Routes

import (
	"net/http"

	Midllewares "github.com/kubitre/blog/Middlewares"
	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
)

type (
	/*RouteSetting - Main route layer for settings all routes*/
	RouteSetting struct {
		APIVersion    string                  // версия api
		Responser     *Midllewares.Responser  // мидлварь
		SecurityLayer *Midllewares.JWTChecker // мидлварь
		Router        *mux.Router             // основной роутер приложения
	}

	/*RouteCRUDs - список роутов до основных операций crud на каждый маршрут*/
	RouteCRUDs struct {
		RouteCreate  string
		RouteUpdate  string
		RouteDelete  string
		RouteFind    string
		RouteFindAll string
	}

	/*IRouter - основной crud всех роутеров*/
	IRouter interface {
		Create(w http.ResponseWriter, r *http.Request)
		Update(w http.ResponseWriter, r *http.Request)
		Remove(w http.ResponseWriter, r *http.Request)
		Find(w http.ResponseWriter, r *http.Request)
		FindAll(w http.ResponseWriter, r *http.Request)
	}

	/*ISetting - интерфейс всех роутеров для быстрого включения\выключения какого обработчика*/
	ISetting interface {
		Setting([]int, *mgo.Database)
		SetupRouterSetting(*RouteSetting)
	}
)

var (
	features = map[int]ISetting{
		0: &ArticleRoute{},
		1: &CommentsRoute{},
		2: &TagRoute{},
		3: &TokenRoute{},
		4: &UserRoute{},
	}
)

/*GetVersion - function for getting version of api*/
func (rs *RouteSetting) GetVersion(w http.ResponseWriter, r *http.Request) {

	rs.Responser.ResponseWithJSON(w, r, http.StatusOK, map[string]string{"version": "0.1"})
}

/*GetStatus - function for getting status on backend*/
func (rs *RouteSetting) GetStatus(w http.ResponseWriter, r *http.Request) {
	rs.Responser.ResponseWithJSON(w, r, http.StatusOK, map[string]string{"current_status": "dev"})
}

/*GetAvailableFormats - function for getting available formats to response type*/
func (rs *RouteSetting) GetAvailableFormats(w http.ResponseWriter, r *http.Request) {
	rs.Responser.ResponseWithJSON(w, r, http.StatusOK, map[string][]string{"available formats": []string{"application/json", "application/xml"}})
}

/*Setting - function for settings database and routes*/
func (rs *RouteSetting) Setting() {

	rs.Router.HandleFunc(rs.APIVersion+"/ver", rs.GetVersion).Methods("GET")
	rs.Router.HandleFunc(rs.APIVersion+"/status", rs.GetStatus).Methods("GET")
	rs.Router.HandleFunc(rs.APIVersion+"/available", rs.GetAvailableFormats).Methods("GET")
	rs.Router.NotFoundHandler = http.HandlerFunc(Midllewares.NotFound)
	rs.Router.MethodNotAllowedHandler = http.HandlerFunc(Midllewares.NotAllowed)
}

/*CreateNewRouter - создание нового роутера*/
func CreateNewRouter(version string) *RouteSetting {
	rs := &RouteSetting{
		Responser: &Midllewares.Responser{
			Error: false,
		},
		APIVersion:    version,
		SecurityLayer: &Midllewares.JWTChecker{},
		Router:        mux.NewRouter(),
	}

	return rs
}

/*StartModeRouters - включение\отключение функционала блога на лету*/
func StartModeRouters(numbersFeatures map[int][]int, routerSetting *RouteSetting, database *mgo.Database) {
	for indexOfFeature, RouteFeatures := range numbersFeatures {
		features[indexOfFeature].SetupRouterSetting(routerSetting)
		features[indexOfFeature].Setting(RouteFeatures, database)
	}

	routerSetting.Setting()
}

/*ConfigureRouterWithFeatures - конфигурация основного роутера с фичами features*/
func (rs *RouteSetting) ConfigureRouterWithFeatures(router IRouter, features []int, routes RouteCRUDs) {
	for _, feature := range features {
		switch feature {
		case 0:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteCreate, router.Create).Methods("POST")
			break
		case 1:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteFind, router.Find).Methods("GET")
			break
		case 2:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteFindAll, router.FindAll).Methods("GET")
			break
		case 3:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteUpdate, router.Update).Methods("PUT")
			break
		case 4:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteDelete, router.Remove).Methods("DELETE")
			break
		}
	}
}
