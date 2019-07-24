package Routes

import (
	"net/http"

	"blog_module/Dao"
	. "blog_module/Middlewares"

	mgo "gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
)

type (
	/*RouteSetting - Main route layer for settings all routes*/
	RouteSetting struct {
		APIVersion    string      // версия api
		Responser     *Responser  // мидлварь
		SecurityLayer *JWTChecker // мидлварь
		Router        *mux.Router // основной роутер приложения
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
		Setting(map[MiddleWare][]Permission, *mgo.Database)
		SetupRouterSetting(*RouteSetting)
	}

	/*Permission - право*/
	Permission int

	/*Features - фича*/
	Features int

	/*MiddleWare - прослойка*/
	MiddleWare int
)

var (
	features = map[Features]ISetting{
		Article: &ArticleRoute{},
		Comment: &CommentsRoute{},
		Tag:     &TagRoute{},
		Token:   &TokenRoute{},
		User:    &UserRoute{},
	}
	lr *LoginRoute = nil
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
func (rs *RouteSetting) Setting(db *mgo.Database) {

	rs.Router.HandleFunc(rs.APIVersion+"/ver", rs.GetVersion).Methods("GET")
	rs.Router.HandleFunc(rs.APIVersion+"/status", rs.GetStatus).Methods("GET")
	rs.Router.HandleFunc(rs.APIVersion+"/available", rs.GetAvailableFormats).Methods("GET")

	rs.Router.NotFoundHandler = http.HandlerFunc(NotFound)
	rs.Router.MethodNotAllowedHandler = http.HandlerFunc(NotAllowed)

	lr := LoginRoute{}
	lr.Setting(nil)
	lr.SetupRouterSetting(rs)
	lr.SetupDaos(db)

	rs.SecurityLayer.ConfiguratingJWTWithSigningMethod(lr.MiddleAuth)

	rs.Router.HandleFunc(rs.APIVersion+lr.Routes.RouteCreate, lr.Authentication).Methods("POST")
	rs.Router.HandleFunc(rs.APIVersion+lr.Routes.RouteDelete, lr.Logout).Methods("DELETE")
}

/*CreateNewRouter - создание нового роутера*/
func CreateNewRouter(version string, database *mgo.Database) *RouteSetting {
	responser := &Responser{
		Error: false,
	}
	rs := &RouteSetting{
		Responser:  responser,
		APIVersion: version,
		SecurityLayer: &JWTChecker{
			Responser: responser,
		},
		Router: mux.NewRouter(),
	}

	rs.SecurityLayer.DaoToken = &Dao.TokenDao{
		Database: database,
	}

	return rs
}

/*StartModeRouters - включение\отключение функционала блога на лету*/
func StartModeRouters(numbersFeatures map[Features]map[MiddleWare][]Permission, routerSetting *RouteSetting, database *mgo.Database) {
	for indexOfFeature, RouteFeatures := range numbersFeatures {
		features[indexOfFeature].SetupRouterSetting(routerSetting)
		features[indexOfFeature].Setting(RouteFeatures, database)
	}

	routerSetting.Setting(database)
}

/*ConfigureMiddlewaresWithFeatures - конфигурация прослоек middleware*/
func (rs *RouteSetting) ConfigureMiddlewaresWithFeatures(router IRouter, middlewares map[MiddleWare][]Permission, routes RouteCRUDs) {
	for index, middleware := range middlewares {
		switch index {
		case Auth:
			rs.ConfigureAuthWithFeatures(router, middleware, routes)
			break
		case Routes:
			rs.ConfigureRoutesWithFeatures(router, middleware, routes)
			break
		}
	}
}

/*ConfigureAuthWithFeatures - конфигурирование прослойки Auth*/
func (rs *RouteSetting) ConfigureAuthWithFeatures(router IRouter, features []Permission, routes RouteCRUDs) {
	for _, feature := range features {
		switch feature {
		case Create:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteCreate, rs.SecurityLayer.JWTCMiddleware(router.Create)).Methods("POST")
			break
		case Read:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteFind, rs.SecurityLayer.JWTCMiddleware(router.Find)).Methods("GET")
			break
		case ReadAll:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteFindAll, rs.SecurityLayer.JWTCMiddleware(router.FindAll)).Methods("GET")
			break
		case Update:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteUpdate, rs.SecurityLayer.JWTCMiddleware(router.Update)).Methods("PUT")
			break
		case Remove:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteDelete, rs.SecurityLayer.JWTCMiddleware(router.Remove)).Methods("DELETE")
			break
		}
	}
}

/*ConfigureRoutesWithFeatures - конфигрурирование роутеров*/
func (rs *RouteSetting) ConfigureRoutesWithFeatures(router IRouter, features []Permission, routes RouteCRUDs) {
	for _, feature := range features {
		switch feature {
		case Create:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteCreate, router.Create).Methods("POST")
			break
		case Read:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteFind, router.Find).Methods("GET")
			break
		case ReadAll:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteFindAll, router.FindAll).Methods("GET")
			break
		case Update:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteUpdate, router.Update).Methods("PUT")
			break
		case Remove:
			rs.Router.HandleFunc(rs.APIVersion+routes.RouteDelete, router.Remove).Methods("DELETE")
			break
		}
	}
}
