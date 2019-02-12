package Routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kubitre/blog/Dao"
	"github.com/kubitre/blog/Models"

	"github.com/gorilla/mux"
)

const (
	apiRouteUsers = "/v1/users"
)

var daoUser = Dao.SettingUser{}

/*UserRoute - Structure for route emdedeed*/
type UserRoute struct {
	ErrorsCounts int32
}

/*CreateNewUser - function for creating new User*/
func (routeSetting *UserRoute) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle create user blog")
	defer r.Body.Close()
	var user Models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid payload!")
		return
	}

	// log.Println("handling user: ", user)

	if err := daoUser.InsertDb(user); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid operation")
		return
	}

	respondWithJSON(w, r, http.StatusOK, "user success created")
}

/*FindUserByID - function for finding User by indentificator*/
func (routeSetting *UserRoute) FindUserByID(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle find user from blog by id")
	params := mux.Vars(r)
	user, err := daoUser.FindByID(params["id"])
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "invalid payload!")
		return
	}

	respondWithJSON(w, r, http.StatusOK, user)
}

/*FindAllUsers - function for finding all Users in database*/
func (routeSetting *UserRoute) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle find all users from blog")
	users, err := daoUser.FindAll()
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "invalid operation!")
		return
	}

	respondWithJSON(w, r, http.StatusOK, users)
}

/*UpdateUserByID - function for updating User by indentificator*/
func (routeSetting *UserRoute) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle update user from blog")
	var user Models.User
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "invalid payload")
		return
	}

	if err := daoUser.Update(user); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "invalid operation")
		return
	}

	respondWithJSON(w, r, http.StatusOK, "user successfully updated!")
}

/*DeleteUserByID - function for remove User by indentificator*/
func (routeSetting *UserRoute) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	// log.Println("handle delete user from blog")
	var user Models.User
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid payload")
		return
	}

	if err := daoUser.Delete(user); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid operation")
		return
	}

	respondWithJSON(w, r, http.StatusOK, "user deleted complete")
}

/*StartSettingRouterUser - function for setting router for articles*/
func StartSettingRouterUser(router *mux.Router) {

	var rout = UserRoute{}

	router.HandleFunc(apiRouteUsers, rout.CreateNewUser).Methods("POST")
	router.HandleFunc(apiRouteUsers, rout.FindAllUsers).Methods("GET")
	router.HandleFunc(apiRouteUsers, rout.FindUserByID).Methods("GET")
	router.HandleFunc(apiRouteUsers, rout.UpdateUserByID).Methods("PUT")
	router.HandleFunc(apiRouteUsers, rout.DeleteUserByID).Methods("DELETE")

	log.Println("routes for users was configurated")
}
