package Routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kubitre/blog/Dao"
	mgo "gopkg.in/mgo.v2"

	"github.com/kubitre/blog/Models"

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

/*CreateNewUser - function for creating new User*/
func (routeSetting *UserRoute) Create(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle create user blog")
	defer r.Body.Close()
	var user Models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}

	// log.Println("handling user: ", user)

	if err := routeSetting.DAO.InsertDb(user); err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid operation",
			"errorCode": err.Error(),
		})
		return
	}

	routeSetting.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, "user success created")
}

/*FindUserByID - function for finding User by indentificator*/
func (routeSetting *UserRoute) Find(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle find user from blog by id")
	params := mux.Vars(r)
	user, err := routeSetting.DAO.FindByID(params["id"])
	if err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}

	routeSetting.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, user)
}

/*FindAllUsers - function for finding all Users in database*/
func (routeSetting *UserRoute) FindAll(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle find all users from blog")
	users, err := routeSetting.DAO.FindAll()
	if err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid operation",
			"errorCode": err.Error(),
		})
		return
	}

	routeSetting.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, users)
}

/*UpdateUserByID - function for updating User by indentificator*/
func (routeSetting *UserRoute) Update(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle update user from blog")
	var user Models.User
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}

	if err := routeSetting.DAO.Update(user); err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid operation",
			"errorCode": err.Error(),
		})
		return
	}

	routeSetting.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, "user successfully updated!")
}

/*DeleteUserByID - function for remove User by indentificator*/
func (routeSetting *UserRoute) Remove(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle delete user from blog")
	var user Models.User
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}

	if err := routeSetting.DAO.Delete(user); err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid operation",
			"errorCode": err.Error(),
		})
		return
	}

	routeSetting.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, "user deleted complete")
}

/*Setting - настройка роутера*/
func (rs *UserRoute) Setting(features []int, db *mgo.Database) {
	rs.Routes = RouteCRUDs{
		RouteCreate:  "/user",
		RouteDelete:  "/users/{id}",
		RouteFind:    "/users/{id}", //  поиск комментария по
		RouteFindAll: "/users",
		RouteUpdate:  "/users/{id}",
	}

	// fmt.Println("Current router: ", *rs)

	var routr IRouter
	routr = rs

	rs.DAO = &Dao.SettingUser{
		Database: db,
	}

	rs.RI.ConfigureRouterWithFeatures(routr.(IRouter), features, rs.Routes)

	log.Println("User route was settinged!")
}

/*SetupRouterSetting - установка главного роутера приложения*/
func (rs *UserRoute) SetupRouterSetting(rS *RouteSetting) {
	rs.RI = rS
}
