package Routes

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEmptyDbArticle(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/articles", nil)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusNoContent, res.Code)

	if body := res.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	req.Header.Add("Content-Type", "application/json")

	ap := ApplicationForTesting{}
	ap.Confugrator()
	rr := httptest.NewRecorder()
	ap.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code: %d. God %d\n", expected, actual)
	}
}
