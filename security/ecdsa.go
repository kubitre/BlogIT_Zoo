package security

import (
	"github.com/jenazads/gojwt"
	"github.com/kubitre/blog/Dao"
	"github.com/kubitre/blog/Models"
	mgo "gopkg.in/mgo.v2"
)

/*ECDSAMiddle - Структура для обработки пользовательских данных*/
type ECDSAMiddle struct {
	Database   *mgo.Database
	GoJWTECDSA *gojwt.Gojwt // шифрование методом ecdsa

}

/*Login - Проверка введённых данных по бд*/
func (ec *ECDSAMiddle) Login(obj Models.Login) (*Models.Token, error) {
	usrInBd, err := ec.CheckEnterUserToBD(obj)

	if err != nil {
		return nil, err
	}

	daoToken := &Dao.TokenDao{
		Database: ec.Database,
	}

	token, err := daoToken.FindToken(usrInBd.ID)
	if err != nil {
		// log.Println("Поиск токена: ", err)
		token, err := ec.GoJWTECDSA.CreateToken(obj.Username)
		tokModel, err := daoToken.CreateNewToken(token, usrInBd.ID)

		return &tokModel, err
	}
	return token, nil
}

func (ec *ECDSAMiddle) CheckEnterUserToBD(obj Models.Login) (*Models.User, error) {
	daoUsers := &Dao.SettingUser{
		Database: ec.Database,
	}

	usr, err := daoUsers.FindByUserCredentials(&obj)
	if err != nil {
		return nil, err
	} else {
		return &usr, nil
	}

}

/*Registration - регистрация нового пользователя*/
func (ec *ECDSAMiddle) Registration(obj *Models.User) error {
	return nil
}

/*Logout - выход из системы*/
func (ec *ECDSAMiddle) Logout(token string) error {
	return nil
}

/*TokenValidation - проверка валидаци токена (распарс токена, проверка времени жизни токена, уровня доступа)*/
func (ec *ECDSAMiddle) TokenValidation(token string) (bool, error) {
	result, _, err := ec.GoJWTECDSA.ValidateToken(token)
	return result, err
}

/*InitInstance - инициализация объекта ECDSA*/
func (ec *ECDSAMiddle) InitInstance(db *mgo.Database) error {
	gojwtO1, err := gojwt.NewGojwtECDSA("kubitre_blog_server", "Access_key", "/keys/priv_key.pem", "/keys/pub_key.pem", "384", 24)

	if err != nil {
		return err
	}
	ec.GoJWTECDSA = gojwtO1

	ec.Database = db
	return nil
}
