package Midllewares

import (
	"net/http"
)

func NotFound(w http.ResponseWriter, r *http.Request) {

	resp := Responser{
		Error: true,
	}

	resp.ResponseWithError(w, r, http.StatusNotFound, map[string]string{
		"Error":  "Not found",
		"Status": "Handled",
	})
}

func NotAllowed(w http.ResponseWriter, r *http.Request) {
	resp := Responser{
		Error: true,
	}

	resp.ResponseWithError(w, r, http.StatusMethodNotAllowed, map[string]string{
		"Error":  "Method not allowed!",
		"Status": "Request failed",
	})
}
