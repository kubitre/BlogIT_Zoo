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
	apiRouteComments = "/comments"
)

var daoComment = Dao.SettingComment{}

/*CommentsRoute - Structure for route emdedeed*/
type CommentsRoute struct {
	ErrorsCounts int32
}

/*CreateNewComment - function for creating new comment*/
func (routeSetting *CommentsRoute) CreateNewComment(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var comment Models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid payload!")
		return
	}
	erra := daoComment.InsertDb(comment)
	if erra != nil {
		respondWithError(w, r, http.StatusInternalServerError, "error create new comment")
		return
	}
	respondWithJSON(w, r, http.StatusOK, comment)
}

/*FindCommentByID - function for finding comment by indentificator*/
func (routeSetting *CommentsRoute) FindCommentByID(w http.ResponseWriter, r *http.Request) {
	var params = mux.Vars(r)
	comment, err := daoComment.FindByID(params["id"])
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "invalid payload")
		return
	}
	respondWithJSON(w, r, http.StatusOK, comment)
}

/*FindAllComments - function for finding all comments in database*/
func (routeSetting *CommentsRoute) FindAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := daoComment.FindAll()
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "invalid operation")
		return
	}
	respondWithJSON(w, r, http.StatusOK, comments)
}

/*UpdateCommentByID - function for updating comment by indentificator*/
func (routeSetting *CommentsRoute) UpdateCommentByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var comment Models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		respondWithError(w, r, http.StatusExpectationFailed, "Invalid payload!")
		return
	}

	if err := daoComment.Update(comment); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid operation")
		return
	}
	respondWithJSON(w, r, http.StatusOK, "success update")
}

/*DeleteCommentByID - function for remove comment by indentificator*/
func (routeSetting *CommentsRoute) DeleteCommentByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var comment Models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		respondWithError(w, r, http.StatusExpectationFailed, "Invalid payload")
		return
	}
	if err := daoComment.Delete(comment); err != nil {
		respondWithError(w, r, http.StatusExpectationFailed, "error delete object")
		return
	}
	respondWithJSON(w, r, http.StatusOK, "success delete object")
}

/*StartSettingRouterComment - function for setting router for articles*/
func StartSettingRouterComment(router *mux.Router) {

	var rout = CommentsRoute{}

	router.HandleFunc(apiRoutearticle, rout.CreateNewComment).Methods("POST")
	router.HandleFunc(apiRoutearticle, rout.FindAllComments).Methods("GET")
	router.HandleFunc(apiRoutearticle, rout.FindCommentByID).Methods("GET")
	router.HandleFunc(apiRoutearticle, rout.UpdateCommentByID).Methods("PUT")
	router.HandleFunc(apiRoutearticle, rout.DeleteCommentByID).Methods("DELETE")

	log.Println("router for comments was configurated")
}
