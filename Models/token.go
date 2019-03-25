package Models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

/*Token - it is structure for access layer of our system*/
type Token struct {
	ID         bson.ObjectId `bson:"_id" json:"-"`               // идентификатор токена в бд
	Value      string        `bson:"maintoken" json:"token"`     // значение токена
	Createdat  time.Time     `bson:"createdat" json:"createdat"` // дата создания
	ValidateTo int64         `bson:"valideto" json:"-"`          // время валидности токена с момента создания
	UserID     bson.ObjectId `bson:"id_user" json:"-"`           // кому принадлежит токен
}

// var signingKey = []byte("secret")

// /*CreateToken - it is function for creating token for accesing layer*/
// func (t *Token) CreateToken(name string) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["name"] = name
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
// 	t.ValidateTo = time.Now().Add(time.Hour * 24).Unix()

// 	tokenString, _ := token.SignedString(signingKey)
// 	t.Value = tokenString
// }

// /*ValidateToken - it is function for validating token in db*/
// func (t *Token) ValidateToken(tok string) {
// 	// var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
// 	// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
// 	// 		return signingKey, nil
// 	// 	},
// 	// 	SigningMethod: jwt.SigningMethodHS256,
// 	// })
// }

// /*RefreshToken - it is function for refreshing token*/
// func (t *Token) RefreshToken(name string) {
// 	t.CreateToken(name)
// }
