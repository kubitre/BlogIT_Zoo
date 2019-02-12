package Routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	apiRoutearticle = "/v1/articles"
)

/*ArticleRoute - Structure for route emdedeed*/
type ArticleRoute struct {
	ErrorsCounts int32
	ApiRoutearticle
}

/*CreateNewArticle - function for creating new article*/
func (routeSetting *ArticleRoute) CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	return
}

/*FindArticleByID - function for finding article by indentificator*/
func (routeSetting *ArticleRoute) FindArticleByID(w http.ResponseWriter, r *http.Request) {
	return
}

/*FindAllArticles - function for finding all articles in database*/
func (routeSetting *ArticleRoute) FindAllArticles(w http.ResponseWriter, r *http.Request) {
	return
}

/*UpdateArticleByID - function for updating article by indentificator*/
func (routeSetting *ArticleRoute) UpdateArticleByID(w http.ResponseWriter, r *http.Request) {
	return
}

/*DeleteArticleByID - function for remove article by indentificator*/
func (routeSetting *ArticleRoute) DeleteArticleByID(w http.ResponseWriter, r *http.Request) {
	return
}

/*StartSettingRouter - function for setting router for articles*/
func StartSettingRouter() (router mux.Router) {

	var rout = ArticleRoute{}
	router = mux.NewRouter()

	router.HandleFunc(apiRoutearticle, rout.CreateNewArticle).Methods("POST")
	router.HandleFunc(apiRoutearticle, rout.FindAllArticles).Methods("GET")
	router.HandleFunc(apiRoutearticle, rout.FindArticleByID).Methods("GET")
	router.HandleFunc(apiRoutearticle, rout.UpdateArticleByID).Methods("UPDATE")
	router.HandleFunc(apiRoutearticle, rout.DeleteArticleByID).Methods("DELETE")

	
}
