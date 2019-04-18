package Routes

import (
	"log"

	"blog_module/Dao"
	"blog_module/Models"

	mgo "gopkg.in/mgo.v2"
)

/*TokenRoute - it is structure which contain info about errors and contain routing handlers*/
type TokenRoute struct {
	Routes RouteCRUDs
	RI     *RouteSetting
	DAO    *Dao.TokenDao

	IRouter
	ISetting
}

/*CreateNewToken - it is function for creating new token*/
func (rs *TokenRoute) CreateNewToken(tok *Models.Token) error {
	return nil
}

/*FindAllToken - it is function for getting all tokens*/
func (rs *TokenRoute) FindAllToken(tok *Models.Token) ([]Models.Token, error) {
	return nil, nil
}

/*FindTokenByID - it is structure for getting token by her id*/
func (rs *TokenRoute) FindTokenByID(id string) (Models.Token, error) {
	return Models.Token{}, nil
}

/*UpdateToken - it is function for refreshing token*/
func (rs *TokenRoute) UpdateToken() (Models.Token, error) {
	return Models.Token{}, nil
}

/*Setting - настройка роутера*/
func (rs *TokenRoute) Setting(middlewares map[MiddleWare][]Permission, db *mgo.Database) {
	rs.Routes = RouteCRUDs{
		RouteCreate: "/token",
		// RouteDelete: "/comments/{id}",
		RouteFind: "/token/{id}", //  поиск комментария по
		// RouteFindAll: "/comments/{id}",
		// RouteUpdate: "/comments/{id}",
	}

	// fmt.Println("Current router: ", *rs)

	rs.DAO = &Dao.TokenDao{
		Database: db,
	}

	var routr IRouter
	routr = rs

	rs.RI.ConfigureMiddlewaresWithFeatures(routr.(IRouter), middlewares, rs.Routes)

	log.Println("Token route was settinged!")
}

// /*StartSettingRouterToken - настройка роутера для токенов*/
// func StartSettingRouterToken(router *mux.Router, routeSetting RouteSetting, JWTMiddle Midllewares.JWTChecker) *mux.Router {
// 	return router
// }

/*SetupRouterSetting - установка главного роутера приложения*/
func (rs *TokenRoute) SetupRouterSetting(rS *RouteSetting) {
	rs.RI = rS
}
