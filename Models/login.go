package Models

/*Login - payload from client app*/
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
