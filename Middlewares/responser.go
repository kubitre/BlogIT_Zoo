package Midllewares

import (
	"encoding/json"
	"net/http"

	"github.com/kubitre/blog/Logs"
)

/*Responser - тип содержащий базовые методы для обработок ошибочных и успешных запросов*/
type Responser struct {
	Tururu string
	Error  bool
}

/*ResponseWithError - ответ с ошибкой*/
func (resp *Responser) ResponseWithError(w http.ResponseWriter, r *http.Request, httpStatus int, payload interface{}) {
	resp.ResponseWithJSON(w, r, httpStatus, payload)
}

/*ResponseWithJSON - ответ в формате json*/
func (resp *Responser) ResponseWithJSON(w http.ResponseWriter, r *http.Request, httpStatus int, payload interface{}) {
	// log := Logger{}
	// log.ConfigTrace(r.Method, r.RequestURI, r.Header.Get("Content-Type"), timestart).PrintTraceRoute()

	if r.Header.Get("Content-Type") != "application/json" {
		Logs.PrintRouteTrace(r, true)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusConflict)
		response, _ := json.Marshal(map[string]string{"error": "you packet in non json format!"})

		w.Write(response)

	} else {
		Logs.PrintRouteTrace(r, resp.Error)

		response, _ := json.Marshal(payload)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpStatus)

		w.Write(response)

	}
}
