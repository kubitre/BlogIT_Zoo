package Routes

import (
	"encoding/json"
	"log"
	"net/http"

	"blog_module/Dao"
	"blog_module/Models"
	"blog_module/security"

	"gopkg.in/mgo.v2"
)

type LoginRoute struct {
	Routes     RouteCRUDs            // основные маршруты до хэндлеров CRUD
	RI         *RouteSetting         // основная сущность маршрута
	TokenDao   *Dao.TokenDao         // DAO слой токенов
	UsersDao   *Dao.SettingUser      // DAO слой пользователей
	MiddleAuth *security.ECDSAMiddle // сервис генерации токена
}

/*Authentication - аутентификацтя по логину и паролю*/
func (rs *LoginRoute) Authentication(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var logpayload Models.Login

	if err := json.NewDecoder(r.Body).Decode(&logpayload); err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"error":     "body packet error",
			"errorCode": err.Error(),
		})
		return
	}

	token, err := rs.MiddleAuth.Login(logpayload)
	if err != nil {
		rs.RI.Responser.ResponseWithError(w, r, http.StatusUnauthorized, map[string]string{
			"error": err.Error(),
		})
		return
	}

	rs.RI.Responser.ResponseWithJSON(w, r, http.StatusOK, map[string]interface{}{
		"token":    token.Value,
		"username": logpayload.Username,
		"userid":   token.UserID,
	})
}

/*Logout - выход из учётки по токену*/
func (rs *LoginRoute) Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("вы успешно вышли из своего аккаунта"))
	w.WriteHeader(http.StatusOK)
}

/*Setting - настройка роутера*/
func (rs *LoginRoute) Setting(db *mgo.Database) {
	rs.Routes = RouteCRUDs{
		RouteCreate: "/auth/login",
		RouteDelete: "/auth/logout",
	}

	log.Println("routes for login was configurated")
}

/*SetupRouterSetting - установка главного роутера приложения*/
func (rs *LoginRoute) SetupRouterSetting(rS *RouteSetting) {
	rs.RI = rS
}

/*SetupDaos - устаовка DAO слоёв в объект LoginRoute*/
func (rs *LoginRoute) SetupDaos(db *mgo.Database) {

	rs.TokenDao = &Dao.TokenDao{
		Database: db,
	}
	rs.UsersDao = &Dao.SettingUser{
		Database: db,
	}

	rs.MiddleAuth = &security.ECDSAMiddle{}
	rs.MiddleAuth.InitInstance(db)
}

/*GetJWTMiddleware - получение мидлвари*/
func (rs *LoginRoute) GetJWTMiddleware() *security.ECDSAMiddle {
	return rs.MiddleAuth
}
