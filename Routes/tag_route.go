package Routes

import "net/http"

const (
	apiRouteTag = "/tags"
)

/*TagRoute - Structure for route emdedeed*/
type TagRoute struct {
	ErrorsCounts int32
}

/*CreateNewTag - function for creating new Tag*/
func (routeSetting *TagRoute) CreateNewTag(w http.ResponseWriter, r *http.Request) {
	return
}

/*FindTagByID - function for finding Tag by indentificator*/
func (routeSetting *TagRoute) FindTagByID(w http.ResponseWriter, r *http.Request) {
	return
}

/*FindAllTags - function for finding all tags in database*/
func (routeSetting *TagRoute) FindAllTags(w http.ResponseWriter, r *http.Request) {
	return
}

/*UpdateTagByID - function for updating tag by indentificator*/
func (routeSetting *TagRoute) UpdateTagByID(w http.ResponseWriter, r *http.Request) {
	return
}

/*DeleteTagByID - function for remove tag by indentificator*/
func (routeSetting *TagRoute) DeleteTagByID(w http.ResponseWriter, r *http.Request) {
	return
}
