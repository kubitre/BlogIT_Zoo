package Routes

import "testing"

func TestEmptyRequestBodyComment_GET(t *testing.T) {
	ker := TestingKernel{
		APIRoute:        "/v1/comments",
		TypeRoute:       "comments",
		AppSettings:     ApplicationForTesting{},
		RequestMethod:   "GET",
		ExpectedCode:    200,
		ExpectedMessage: "null"}
	ker.AppSettings.ConfigureDbConnection()
	ker.RunTime(t, nil)
}
