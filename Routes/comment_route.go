package Routes

import "net/http"

const (
	apiRouteComments = "/comments"
)

/*CommentsRoute - Structure for route emdedeed*/
type CommentsRoute struct {
	ErrorsCounts int32
}

/*CreateNewComment - function for creating new comment*/
func (routeSetting *CommentsRoute) CreateNewComment(w http.ResponseWriter, r *http.Request) {
	return
}

/*FindCommentByID - function for finding comment by indentificator*/
func (routeSetting *CommentsRoute) FindCommentByID(w http.ResponseWriter, r *http.Request) {
	return
}

/*FindAllComments - function for finding all comments in database*/
func (routeSetting *CommentsRoute) FindAllComments(w http.ResponseWriter, r *http.Request) {
	return
}

/*UpdateCommentByID - function for updating comment by indentificator*/
func (routeSetting *CommentsRoute) UpdateCommentByID(w http.ResponseWriter, r *http.Request) {
	return
}

/*DeleteCommentByID - function for remove comment by indentificator*/
func (routeSetting *ArticleRoute) DeleteCommentByID(w http.ResponseWriter, r *http.Request) {
	return
}
