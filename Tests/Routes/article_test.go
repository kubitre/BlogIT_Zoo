package Routes

import "testing"

func TestEmptyDbArticle_GET(t *testing.T) {
	ker := TestingKernel{
		APIRoute:        "/v1/articles",
		TypeRoute:       "articles",
		AppSettings:     ApplicationForTesting{},
		RequestMethod:   "GET",
		ExpectedCode:    204,
		ExpectedMessage: "null"}
	ker.AppSettings.ConfigureDbConnection()
	ker.RunTime(t, nil)
}

func TestArticleEmptyCreate_POST(t *testing.T) {
	ker := TestingKernel{
		APIRoute:        "/v1/articles",
		TypeRoute:       "articles",
		AppSettings:     ApplicationForTesting{},
		RequestMethod:   "POST",
		ExpectedCode:    500,
		ExpectedMessage: `{"error": "Not insert to database! Please contact with administration"}`,
	}
	ker.AppSettings.ConfigureDbConnection()
	ker.RunTime(t, []byte(`
		{
			"name": "test", 
			"description": "test description",
			"tags": null,
			"author": null}`))
}
