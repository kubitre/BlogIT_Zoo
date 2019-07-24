package Routes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

/*TestingKernel -  it is structure which contain route which need test and setting for route*/
type TestingKernel struct {
	APIRoute        string
	TypeRoute       string
	AppSettings     ApplicationForTesting
	RequestMethod   string
	ExpectedCode    int
	ExpectedMessage string
}

/*RunTime - it is function for start test explorer*/
func (testSt *TestingKernel) RunTime(t *testing.T, body []byte) {
	req, _ := http.NewRequest(testSt.RequestMethod, testSt.APIRoute, bytes.NewBuffer(body))

	router, _ := testSt.AppSettings.Configurator(testSt.TypeRoute)

	res := executeRequest(router, req)
	checkResponseCode(t, testSt.ExpectedCode, res.Code)

	if body := res.Body.String(); body != testSt.ExpectedMessage {
		t.Errorf("Expected an %s. Got %s", testSt.ExpectedMessage, body)
	}
}

func executeRequest(handl *mux.Router, req *http.Request) *httptest.ResponseRecorder {
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handl.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code: %d. God %d\n", expected, actual)
	}
}
