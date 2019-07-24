package Midllewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"blog_module/security"

	"blog_module/Dao"
)

/*JWTChecker - прослойка для проверки токена*/
type JWTChecker struct {
	JWTMiddleWare *security.ECDSAMiddle
	DaoToken      *Dao.TokenDao
	Responser     *Responser
}

/*JWTCMiddleware - прослойка проверки токена*/
func (tjwt *JWTChecker) JWTCMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: check Token authentification
		tokenHeader := r.Header.Get("Authorization")
		tokenValue := strings.Split(tokenHeader, " ")[1]

		// log.Println("Start authentification check, auth token: ", tokenHeader)

		fmt.Println("Token value: ", tokenValue)

		// 1 STEP: FInd Token in DAO Token layer
		tokenFromDb, err := tjwt.DaoToken.FindTokenByValue(tokenValue)

		if err != nil {
			fmt.Println("Has error find token in database. Error code: ", err.Error())
			tjwt.Responser.ResponseWithError(w, r, http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
				"trace": "jwtmidleware_checker",
			})

			return
		}
		fmt.Println("Token from db: ", tokenFromDb)

		// 2 STEP: validation duration time and created time of token

		// 3 STEP: Checking validatioon of enter token
		result, err := tjwt.JWTMiddleWare.TokenValidation(tokenValue)
		if err != nil && !result {
			tjwt.Responser.ResponseWithError(w, r, http.StatusUnauthorized, map[string]string{
				"error": err.Error(),
			})
		}

		log.Println("Authorized completed!")

		next.ServeHTTP(w, r)
	})
}

/*ConfiguratingJWTWithSigningMethod - конфигурация мидлвари на чек токена, с методом подписи*/
func (tjwt *JWTChecker) ConfiguratingJWTWithSigningMethod(jwtmiddleware *security.ECDSAMiddle) {
	tjwt.JWTMiddleWare = jwtmiddleware
}
