package Models

import "time"

/*Token - it is structure for access layer of our system*/
type Token struct {
	value      string
	createdat  time.Time
	validateTo time.Time
}

/*CreateToken - it is function for creating token for accesing layer*/
func (t *Token) CreateToken() {

}

/*ValidateToken - it is function for validating token in db*/
func (t *Token) ValidateToken(tok string) {

}

/*RefreshToken - it is function for refreshing token*/
func (t *Token) RefreshToken(tok string) {

}
