package Routes

import (
	"encoding/json"
	"log"
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
	ErrorsCounts int32
}

/*CreateNewTag - function for creating new Tag*/
func (routeSetting *TagRoute) CreateNewTag(w http.ResponseWriter, r *http.Request) {
	var tag Models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		r.Body.Close()
		respondWithError(w, r, http.StatusInternalServerError, "Invalid payload")
		return
	}
	if err := daoTag.InsertDb(tag); err != nil {
		r.Body.Close()
		respondWithError(w, r, http.StatusInternalServerError, "invalid insert into db")
		return
	}

	r.Body.Close()
	respondWithJSON(w, r, http.StatusOK, tag)
}

/*FindTagByID - function for finding Tag by indentificator*/
func (routeSetting *TagRoute) FindTagByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tag, err := daoTag.FindByID(params["id"])
	if err != nil {
		respondWithError(w, r, http.StatusExpectationFailed, "Invalid payload!")
		return
	}
	respondWithJSON(w, r, http.StatusOK, tag)
}

/*FindAllTags - function for finding all tags in database*/
func (routeSetting *TagRoute) FindAllTags(w http.ResponseWriter, r *http.Request) {
	tags, err := daoTag.FindAll()
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "invalid operations")
		return
	}
	respondWithJSON(w, r, http.StatusOK, tags)
}

/*UpdateTagByID - function for updating tag by indentificator*/
func (routeSetting *TagRoute) UpdateTagByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var tag Models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid payload")
		return
	}

	if err := daoTag.Update(tag); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid operations")
		return
	}
	respondWithJSON(w, r, http.StatusOK, "success update object")
}

/*DeleteTagByID - function for remove tag by indentificator*/
func (routeSetting *TagRoute) DeleteTagByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var tag Models.Tag
	if err := json.NewDecoder(r.Body).Decode(&tag); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid payload!")
		return
	}
	if err := daoTag.Delete(tag); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid operation")
		return
	}

	respondWithJSON(w, r, http.StatusOK, "delete was complete")
}

/*StartSettingRouterTag - function for setting router for articles*/
func StartSettingRouterTag(router *mux.Router) {

	var rout = TagRoute{}

	router.HandleFunc(apiRouteTag, rout.CreateNewTag).Methods("POST")
	router.HandleFunc(apiRouteTag, rout.FindAllTags).Methods("GET")
	router.HandleFunc(apiRouteTag, rout.FindTagByID).Methods("GET")
	router.HandleFunc(apiRouteTag, rout.UpdateTagByID).Methods("PUT")
	router.HandleFunc(apiRouteTag, rout.DeleteTagByID).Methods("DELETE")

	log.Println("routes for tags was configurated")
}
