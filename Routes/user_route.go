package Routes

import (
	"encoding/json"
	"net/http"

	"github.com/kubitre/blog/Dao"

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
func (routeSetting *UserRoute) CreateNewUser(w http.ResponseWriter, r *http.Request) {
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
func (routeSetting *UserRoute) FindUserByID(w http.ResponseWriter, r *http.Request) {
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
func (routeSetting *UserRoute) FindAllUsers(w http.ResponseWriter, r *http.Request) {
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
func (routeSetting *UserRoute) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
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
func (routeSetting *UserRoute) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
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
func (rs *UserRoute) Setting(features []int) {
}

// /*StartSettingRouterUser - function for setting router for articles*/
// func StartSettingRouterUser(router *mux.Router, routSetting RouteSetting, JWTMiddle Midllewares.JWTChecker) {

// 	// var rout = UserRoute{
// 	// 	Setting: routSetting,
// 	// }

// 	// router.HandleFunc(apiRouteUsers, rout.CreateNewUser).Methods("POST")
// 	// router.HandleFunc(apiRouteUsers, rout.FindAllUsers).Methods("GET")
// 	// router.HandleFunc(apiRouteUsers, rout.FindUserByID).Methods("GET")
// 	// router.HandleFunc(apiRouteUsers, rout.UpdateUserByID).Methods("PUT")
// 	// router.HandleFunc(apiRouteUsers, rout.DeleteUserByID).Methods("DELETE")

// 	log.Println("routes for users was configurated")
// }

/*SetupRouterSetting - установка главного роутера приложения*/
func (rs *UserRoute) SetupRouterSetting(rS *RouteSetting) {
	rs.RI = rS
}
