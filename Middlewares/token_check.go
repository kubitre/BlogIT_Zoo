package Midllewares

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

/*JWTChecker - прослойка для проверки токена*/
type JWTChecker struct {
	jwtMiddleWare jwtmiddleware.JWTMiddleware
}

/*Checker - прослойка с проверкой токена*/
func (tjwt *JWTChecker) Checker(signKey string) jwtmiddleware.JWTMiddleware {
	tjwt.jwtMiddleWare = *jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(signKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return tjwt.jwtMiddleWare
}
