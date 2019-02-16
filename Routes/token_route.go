package Routes

import "github.com/kubitre/blog/Models"

/*TokenRoute - it is structure which contain info about errors and contain routing handlers*/
type TokenRoute struct {
	AmountErrors int32
}

/*CreateNewToken - it is function for creating new token*/
func (t *TokenRoute) CreateNewToken(tok *Models.Token) error {
	return nil
}

/*FindAllToken - it is function for getting all tokens*/
func (t *TokenRoute) FindAllToken(tok *Models.Token) ([]Models.Token, error) {
	return nil, nil
}

/*FindTokenByID - it is structure for getting token by her id*/
func (t *TokenRoute) FindTokenByID(id string) (Models.Token, error) {
	return nil

/*UpdateToken - it is function for refreshing token*/
func (t *TokenRoute) UpdateToken() (Models.Token, error) {
	return nil
}
