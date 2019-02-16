package Routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kubitre/blog/Config"
)

const (
	addr = ":9512"
)

/*respondWithError - function for send error*/
func respondWithError(w http.ResponseWriter, r *http.Request, code int, msg string) {
	respondWithJSON(w, r, code, map[string]string{"error": msg})
}

/*respondWithJSON - function for send packet to response in json type*/
func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	if r.Header.Get("Content-Type") != "application/json" {
		// Logs.PrintRouteTrace(r, true)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		respons, _ := json.Marshal(map[string]string{"error": "you packet in not json format! Please check you packet"})
		w.Write(respons)
	} else {
		// Logs.PrintRouteTrace(r, false)
		response, _ := json.Marshal(payload)
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(code)
		w.Write(response)
	}
}

/*RouteSetting - Main route layer for settings all routes*/
type RouteSetting struct {
	apiVersion string
}

const (
	apiRouteMain = "/v"
	version      = "1"
)

/*GetVersion - function for getting version of api*/
func (routesetting *RouteSetting) GetVersion(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, r, http.StatusOK, map[string]string{"version": "0.1"})
}

/*GetStatus - function for getting status on backend*/
func (routesetting *RouteSetting) GetStatus(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, r, http.StatusOK, map[string]string{"current_status": "dev"})
}

/*GetAvailableFormats - function for getting available formats to response type*/
func (routsetting *RouteSetting) GetAvailableFormats(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, r, http.StatusOK, map[string][]string{"available formats": []string{"application/json", "application/xml"}})
}

/*StartSettingRoutes - function for settings database and routes*/
func StartSettingRoutes(config Config.Configuration, flagTest bool) *mux.Router {
	router := mux.NewRouter()
	routSetting := RouteSetting{}

	StartSettingRouterArticle(router)
	StartSettingRouterComment(router)
	StartSettingRouterTag(router)
	StartSettingRouterUser(router)

	router.HandleFunc(apiRouteMain+version+"/ver", routSetting.GetVersion).Methods("GET")
	router.HandleFunc(apiRouteMain+version+"/status", routSetting.GetStatus).Methods("GET")
	router.HandleFunc(apiRouteMain+version+"/available", routSetting.GetAvailableFormats).Methods("GET")

	if flagTest {
		return router
	}

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}

	return router
}
