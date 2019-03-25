package Midllewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/kubitre/blog/Dao"
)

/*JWTChecker - прослойка для проверки токена*/
type JWTChecker struct {
	jwtMiddleWare jwtmiddleware.JWTMiddleware
	DaoToken      *Dao.TokenDao
}

/*CheckerLayerConfiguration - прослойка с проверкой токена*/
func (tjwt *JWTChecker) CheckerLayerConfiguration(signKey string) jwtmiddleware.JWTMiddleware {
	tjwt.jwtMiddleWare = *jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(signKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return tjwt.jwtMiddleWare
}

/*JWTCMiddleware - прослойка проверки токена*/
func (tjwt *JWTChecker) JWTCMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: check Token authentification
		tokenHeader := r.Header.Get("Authorization")
		tokenValue := strings.Split(tokenHeader, " ")[1]

		log.Println("Start authentification check, auth token: ", tokenHeader)

		fmt.Println("Token value: ", tokenValue)

		// 1 STEP: FInd Token in DAO Token layer
		tokenFromDb, err := tjwt.DaoToken.FindTokenByValue(tokenValue)
		if err != nil {
			fmt.Println("Has error find token in database. Error code: ", err.Error())
		}
		fmt.Println("Token from db: ", tokenFromDb)

		// 2 STEP: Checking validatioon of enter token

		next.ServeHTTP(w, r)
	})
}
