package Routes

import "net/http"

const (
	apiRouteUsers = "/users"
)

/*UserRoute - Structure for route emdedeed*/
type UserRoute struct {
	ErrorsCounts int32
}

/*CreateNewUser - function for creating new User*/
func (routeSetting *UserRoute) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	return
}

/*FindUserByID - function for finding User by indentificator*/
func (routeSetting *UserRoute) FindUserByID(w http.ResponseWriter, r *http.Request) {
	return
}

/*FindAllUsers - function for finding all Users in database*/
func (routeSetting *UserRoute) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	return
}

/*UpdateUserByID - function for updating User by indentificator*/
func (routeSetting *UserRoute) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	return
}

/*DeleteUserByID - function for remove User by indentificator*/
func (routeSetting *UserRoute) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	return
}
