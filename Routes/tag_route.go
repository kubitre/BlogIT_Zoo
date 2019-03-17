package Routes

import (
	"encoding/json"
	"net/http"

	"github.com/kubitre/blog/Dao"

	"github.com/gorilla/mux"
	"github.com/kubitre/blog/Models"
)

const (
	apiRouteTag = "/v1/tags"
)

var daoTag = Dao.SettingTag{}

/*TagRoute - Structure for route emdedeed*/
type TagRoute struct {
	Routes RouteCRUDs
	RI     *RouteSetting
	DAO    *Dao.SettingTag

	IRouter
	ISetting
}

/*Create - function for creating new Tag*/
func (routeSetting *TagRoute) Create(w http.ResponseWriter, r *http.Request) {
	var tag Models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		r.Body.Close()
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload",
			"errorCode": err.Error(),
		})
		return
	}
	tagInserted, err := daoTag.InsertDb(tag)
	if err != nil {
		r.Body.Close()
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{"error": "invalid insert into db", "errorCode": err.Error()})
		return
	}

	r.Body.Close()
	routeSetting.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, tagInserted)
}

/*Find - function for finding Tag by indentificator*/
func (routeSetting *TagRoute) Find(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag, err := daoTag.FindByID(params["id"])
	if err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusExpectationFailed, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}
	routeSetting.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, tag)
}

/*FindAll - function for finding all tags in database*/
func (routeSetting *TagRoute) FindAll(w http.ResponseWriter, r *http.Request) {
	tags, err := daoTag.FindAll()
	if err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "invalid operations",
			"errorCode": err.Error(),
		})
		return
	}
	routeSetting.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, tags)
}

/*Update - function for updating tag by indentificator*/
func (routeSetting *TagRoute) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var tag Models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}

	if err := daoTag.Update(tag); err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "invalid operations",
			"errorCode": err.Error(),
		})
		return
	}
	routeSetting.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, "success update object")
}

/*Remove - function for remove tag by indentificator*/
func (routeSetting *TagRoute) Remove(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var tag Models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}
	if err := daoTag.Delete(tag); err != nil {
		routeSetting.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "invalid operations",
			"errorCode": err.Error(),
		})
		return
	}

	routeSetting.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, map[string]string{
		"status": "delete was complete",
		"code":   "test",
	})
}

// /*StartSettingRouterTag - function for setting router for articles*/
// func StartSettingRouterTag(router *mux.Router, routSetting RouteSetting, JWTMiddle Midllewares.JWTChecker) {

// 	var rout = TagRoute{}

// 	router.HandleFunc(apiRouteTag, rout.CreateNewTag).Methods("POST")
// 	router.HandleFunc(apiRouteTag, rout.FindAllTags).Methods("GET")
// 	router.HandleFunc(apiRouteTag, rout.FindTagByID).Methods("GET")
// 	router.HandleFunc(apiRouteTag, rout.UpdateTagByID).Methods("PUT")
// 	router.HandleFunc(apiRouteTag, rout.DeleteTagByID).Methods("DELETE")

// 	log.Println("routes for tags was configurated")
// }

func (rs *TagRoute) Setting(features []int) {
}

/*SetupRouterSetting - установка главного роутера приложения*/
func (rs *TagRoute) SetupRouterSetting(rS *RouteSetting) {
	rs.RI = rS
}
