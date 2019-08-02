package Routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"blog_module/Dao"
	"blog_module/Models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

/*CommentsRoute - Structure for route emdedeed*/
type CommentsRoute struct {
	Routes RouteCRUDs
	RI     *RouteSetting
	DAO    *Dao.SettingComment

	IRouter
	ISetting
}

/*Create - function for creating new comment*/
func (rs *CommentsRoute) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var comment Models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "invalid paylaod",
			"errorCode": err.Error(),
		})
		return
	}
	newcomment, erra := rs.DAO.InsertDb(comment)
	if erra != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "does not created new comment!",
			"errorCode": erra.Error(),
		})
		return
	}
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, newcomment)
}

/*Find - function for finding comment by indentificator*/
func (rs *CommentsRoute) Find(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	fmt.Println("[COMMENT_ROUTE]: start fetching")
	comment, err := rs.DAO.FindByID(params["id"])
	if err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "invalid payload",
			"errorCode": err.Error(),
		})
		return
	}
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, comment)
}

/*FindAll - function for finding all comments in database*/
func (rs *CommentsRoute) FindAll(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	fmt.Println("[COMMENT_ROUTE]: start fetching")
	comments, err := rs.DAO.FindAll(params["id"])
	if err != nil {

		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "invalid operation",
			"errorCode": err.Error(),
		})
		return
	}
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, comments)
}

/*Update - function for updating comment by indentificator*/
func (rs *CommentsRoute) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var comment Models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusExpectationFailed, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}

	if err := rs.DAO.Update(comment); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid operation",
			"errorCode": err.Error(),
		})
		return
	}
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, map[string]string{
		"status": "success update",
		"code":   "test",
	})
}

/*Remove - function for remove comment by indentificator*/
func (rs *CommentsRoute) Remove(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var comment Models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusExpectationFailed, map[string]string{
			"error":     "Invalid payload",
			"errorCode": err.Error(),
		})
		return
	}
	if err := rs.DAO.Delete(comment); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusExpectationFailed, map[string]string{
			"error":     "error delete object",
			"errorCode": err.Error(),
		})
		return
	}
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, map[string]string{
		"status": "success delete object",
		"code":   "test",
	})
}

/*Setting - настроечный интерфейс*/
func (rs *CommentsRoute) Setting(middlewares map[MiddleWare][]Permission, db *mgo.Database) {
	rs.Routes = RouteCRUDs{
		RouteCreate:  "/comment",
		RouteDelete:  "/comments/{id}",
		RouteFind:    "/comments/{id}", //  поиск комментария по
		RouteFindAll: "/commentss/{id}",
		RouteUpdate:  "/comments/{id}",
	}

	rs.DAO = &Dao.SettingComment{
		Database: db,
	}

	var routr IRouter
	routr = rs

	rs.RI.ConfigureMiddlewaresWithFeatures(routr.(IRouter), middlewares, rs.Routes)

	log.Println("Comment route was settinged!")
}

/*SetupRouterSetting - установка главного роутера приложения*/
func (rs *CommentsRoute) SetupRouterSetting(rS *RouteSetting) {
	rs.RI = rS
}
