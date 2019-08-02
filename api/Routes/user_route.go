package Routes

import (
	"encoding/json"
	"log"
	"net/http"

	"blog_module/Dao"

	"gopkg.in/mgo.v2"

	"blog_module/Models"

	"github.com/gorilla/mux"
)

/*UserRoute - Structure for route emdedeed*/
type UserRoute struct {
	Routes RouteCRUDs
	RI     *RouteSetting
	DAO    *Dao.SettingUser

	IRouter
	ISetting
}

/*Create - function for creating new User*/
func (rs *UserRoute) Create(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle create user blog")
	defer r.Body.Close()
	var user Models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}

	if err := rs.DAO.InsertDb(user); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid operation",
			"errorCode": err.Error(),
		})
		return
	}

	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, "user success created")
}

/*Find - function for finding User by indentificator*/
func (rs *UserRoute) Find(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle find user from blog by id")
	params := mux.Vars(r)
	user, err := rs.DAO.FindByID(params["id"])
	if err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}

	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, user)
}

/*FindAll - function for finding all Users in database*/
func (rs *UserRoute) FindAll(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle find all users from blog")
	users, err := rs.DAO.FindAll()
	if err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid operation",
			"errorCode": err.Error(),
		})
		return
	}

	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, users)
}

/*Update - function for updating User by indentificator*/
func (rs *UserRoute) Update(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle update user from blog")
	var user Models.User
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}

	if err := rs.DAO.Update(user); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid operation",
			"errorCode": err.Error(),
		})
		return
	}

	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, "user successfully updated!")
}

/*Remove - function for remove User by indentificator*/
func (rs *UserRoute) Remove(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle delete user from blog")
	var user Models.User
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}

	if err := rs.DAO.Delete(user); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid operation",
			"errorCode": err.Error(),
		})
		return
	}

	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, "user deleted complete")
}

/*Setting - настройка роутера*/
func (rs *UserRoute) Setting(middlewares map[MiddleWare][]Permission, db *mgo.Database) {
	rs.Routes = RouteCRUDs{
		RouteCreate:  "/user",
		RouteDelete:  "/users/{id}",
		RouteFind:    "/users/{id}", //  поиск комментария по
		RouteFindAll: "/users",
		RouteUpdate:  "/users/{id}",
	}

	var routr IRouter
	routr = rs

	rs.DAO = &Dao.SettingUser{
		Database: db,
	}

	rs.RI.ConfigureMiddlewaresWithFeatures(routr.(IRouter), middlewares, rs.Routes)

	log.Println("User route was settinged!")
}

/*SetupRouterSetting - установка главного роутера приложения*/
func (rs *UserRoute) SetupRouterSetting(rS *RouteSetting) {
	rs.RI = rS
}
