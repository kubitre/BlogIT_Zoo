package Routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kubitre/blog/Dao"
	"github.com/kubitre/blog/Models"
)

/*ArticleRoute - Structure for route emdedeed*/
type ArticleRoute struct {
	Routes RouteCRUDs          // основные маршруты до хэндлеров CRUD
	RI     *RouteSetting       // основная сущность маршрута
	DAO    *Dao.SettingArticle // DAO слой

	IRouter // основные хэндлеры маршрута
	ISetting
}

/*CreateArticle - function for creating new article*/
func (rs *ArticleRoute) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var article Models.Article

	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusExpectationFailed, map[string]string{
			"error":     "Invalid request payload!",
			"errorCode": err.Error(),
		})
		return
	}
	artic, err := rs.DAO.InsertDb(article)
	if err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Not insert to database! Please contact with adminitstration",
			"errorCode": err.Error(),
		})
		return
	}
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, artic)
}

/*Find - function for finding article by indentificator*/
func (rs *ArticleRoute) Find(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	artic, err := rs.DAO.FindByID(params["id"])
	if err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid indentificator or not exist! please check you id",
			"errorCode": err.Error(),
		})
		return
	}
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, artic)
}

/*FindAll - function for finding all articles in database*/
func (rs *ArticleRoute) FindAll(w http.ResponseWriter, r *http.Request) {
	articles, err := rs.DAO.FindAll()
	if err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "invalid payload",
			"errorCode": err.Error(),
		})
		return
	}
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusNoContent, articles)
}

/*Update - function for updating article by indentificator*/
func (rs *ArticleRoute) Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var artic Models.Article
	if err := json.NewDecoder(r.Body).Decode(&artic); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid payload!",
			"errorCode": err.Error(),
		})
		return
	}
	if err := rs.DAO.Update(artic); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid update!",
			"errorCode": err.Error(),
		})
		return
	}
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, map[string]string{
		"status": "article was updated!",
		"code":   "test",
	})
}

/*Remove - function for remove article by indentificator*/
func (rs *ArticleRoute) Remove(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var artic Models.Article
	if err := json.NewDecoder(r.Body).Decode(&artic); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "invalid payload",
			"errorCode": err.Error(),
		})
		return
	}
	if err := rs.DAO.Delete(artic); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "Invalid delete object",
			"errorCode": err.Error(),
		})
		return
	}
	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, "delete was complete")
}

/*Setting - function for setting router for articles*/
func (rs *ArticleRoute) Setting(fetures []int) {
	rs.Routes = RouteCRUDs{
		RouteCreate:  "/article",
		RouteDelete:  "/articles/{id}",
		RouteFind:    "/articles/{id}",
		RouteFindAll: "/articles",
		RouteUpdate:  "/articles/{id}",
	}

	// fmt.Println("Current router: ", *rs)

	var routr IRouter
	routr = rs

	rs.RI.ConfigureRouterWithFeatures(routr.(IRouter), fetures, rs.Routes)

	log.Println("routes for articles was configurated")
}

/*SetupRouterSetting - установка главного роутера приложения*/
func (rs *ArticleRoute) SetupRouterSetting(rS *RouteSetting) {
	rs.RI = rS
}
