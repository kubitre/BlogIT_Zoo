package Routes

import (
	"encoding/json"
	"log"
	"net/http"

	"blog_module/Dao"

	mgo "gopkg.in/mgo.v2"

	"blog_module/Models"

	"github.com/gorilla/mux"
)

/*TagRoute - Structure for route emdedeed*/
type TagRoute struct {
	Routes RouteCRUDs
	RI     *RouteSetting
	DAO    *Dao.SettingTag

	IRouter
	ISetting
}

/*Create - function for creating new Tag*/
func (rs *TagRoute) Create(w http.ResponseWriter, r *http.Request) {
	var tag Models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		r.Body.Close()
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload",
			"errorCode": err.Error(),
		})
		return
	}
	tagInserted, err := rs.DAO.InsertDb(tag)
	if err != nil {
		r.Body.Close()
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{"error": "invalid insert into db", "errorCode": err.Error()})
		return
	}

	r.Body.Close()
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, tagInserted)
}

/*Find - function for finding Tag by indentificator*/
func (rs *TagRoute) Find(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag, err := rs.DAO.FindByID(params["id"])
	if err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusExpectationFailed, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, tag)
}

/*FindAll - function for finding all tags in database*/
func (rs *TagRoute) FindAll(w http.ResponseWriter, r *http.Request) {
	tags, err := rs.DAO.FindAll()
	if err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "invalid operations",
			"errorCode": err.Error(),
		})
		return
	}
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, tags)
}

/*Update - function for updating tag by indentificator*/
func (rs *TagRoute) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var tag Models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}

	if err := rs.DAO.Update(tag); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "invalid operations",
			"errorCode": err.Error(),
		})
		return
	}
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, "success update object")
}

/*Remove - function for remove tag by indentificator*/
func (rs *TagRoute) Remove(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var tag Models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}
	if err := rs.DAO.Delete(tag); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "invalid operations",
			"errorCode": err.Error(),
		})
		return
	}

	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, map[string]string{
		"status": "delete was complete",
		"code":   "test",
	})
}

// /*StartSettingRouterTag - function for setting router for articles*/
// func StartSettingRouterTag(router *mux.Router, routSetting rs, JWTMiddle Midllewares.JWTChecker) {

// 	var rout = TagRoute{}

// 	router.HandleFunc(apiRouteTag, rout.CreateNewTag).Methods("POST")
// 	router.HandleFunc(apiRouteTag, rout.FindAllTags).Methods("GET")
// 	router.HandleFunc(apiRouteTag, rout.FindTagByID).Methods("GET")
// 	router.HandleFunc(apiRouteTag, rout.UpdateTagByID).Methods("PUT")
// 	router.HandleFunc(apiRouteTag, rout.DeleteTagByID).Methods("DELETE")

// 	log.Println("routes for tags was configurated")
// }

func (rs *TagRoute) Setting(middlewares map[MiddleWare][]Permission, db *mgo.Database) {
	rs.Routes = RouteCRUDs{
		RouteCreate:  "/tag",
		RouteDelete:  "/tag/{id}",
		RouteFind:    "/tags/{id}", //  поиск комментария по
		RouteFindAll: "/tags",
		RouteUpdate:  "/tags/{id}",
	}

	// fmt.Println("Current router: ", *rs)

	rs.DAO = &Dao.SettingTag{
		Database: db,
	}

	var routr IRouter
	routr = rs

	rs.RI.ConfigureMiddlewaresWithFeatures(routr.(IRouter), middlewares, rs.Routes)

	log.Println("Tag route was settinged!")
}

/*SetupRouterSetting - установка главного роутера приложения*/
func (rs *TagRoute) SetupRouterSetting(rS *RouteSetting) {
	rs.RI = rS
}
