package Routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kubitre/blog/Dao"
	"github.com/kubitre/blog/Models"
)

const (
	apiRoutearticle = "/v1/articles"
)

var dao = Dao.SettingArticle{}

/*ArticleRoute - Structure for route emdedeed*/
type ArticleRoute struct {
	ErrorsCounts int32
}

/*CreateNewArticle - function for creating new article*/
func (routeSetting *ArticleRoute) CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var article Models.Article

	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		respondWithError(w, r, http.StatusExpectationFailed, "Invalid request payload!")
		return
	}
	artic, err := dao.InsertDb(article)
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Not insert to database! Please contact with adminitstration")
		return
	}
	respondWithJSON(w, r, http.StatusOK, artic)
}

/*FindArticleByID - function for finding article by indentificator*/
func (routeSetting *ArticleRoute) FindArticleByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	artic, err := dao.FindByID(params["id"])
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid indentificator or not exist! please check you id")
		return
	}
	respondWithJSON(w, r, http.StatusOK, artic)
}

/*FindAllArticles - function for finding all articles in database*/
func (routeSetting *ArticleRoute) FindAllArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := dao.FindAll()
	if err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Internal server error")
		return
	}
	respondWithJSON(w, r, http.StatusNoContent, articles)
}

/*UpdateArticleByID - function for updating article by indentificator*/
func (routeSetting *ArticleRoute) UpdateArticleByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var artic Models.Article
	if err := json.NewDecoder(r.Body).Decode(&artic); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid payload!")
		return
	}
	if err := dao.Update(artic); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid update!")
		return
	}
	respondWithJSON(w, r, http.StatusOK, "article was updated!")
}

/*DeleteArticleByID - function for remove article by indentificator*/
func (routeSetting *ArticleRoute) DeleteArticleByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var artic Models.Article
	if err := json.NewDecoder(r.Body).Decode(&artic); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "invalid payload")
		return
	}
	if err := dao.Delete(artic); err != nil {
		respondWithError(w, r, http.StatusInternalServerError, "Invalid delete object")
		return
	}
	respondWithJSON(w, r, http.StatusOK, "delete was complete")
}

/*StartSettingRouterArticle - function for setting router for articles*/
func StartSettingRouterArticle(router *mux.Router) {

	var rout = ArticleRoute{}

	router.HandleFunc(apiRoutearticle, rout.CreateNewArticle).Methods("POST")
	router.HandleFunc(apiRoutearticle, rout.FindAllArticles).Methods("GET")
	router.HandleFunc(apiRoutearticle, rout.FindArticleByID).Methods("GET")
	router.HandleFunc(apiRoutearticle, rout.UpdateArticleByID).Methods("PUT")
	router.HandleFunc(apiRoutearticle, rout.DeleteArticleByID).Methods("DELETE")

	log.Println("routes for articles was configurated")
}
