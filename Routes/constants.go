package Routes

/*MRouter - получение номера мдилвари роутеров*/
func MRouter() MiddleWare {
	return Routes
}

/*MLogger - получение номера мидлвари логера*/
func MLogger() MiddleWare {
	return Logger
}

/*MAuth - получение номера мидлвари аутентификаии*/
func MAuth() MiddleWare {
	return Auth
}

/*Cre - получение права на запись*/
func Cre() Permission { return Create }

/*Rea - получение права на чтение по идентификатору*/
func Rea() Permission { return Read }

/*ReA - получение права на чтение по лимиту и оффсету*/
func ReA() Permission { return ReadAll }

/*Upd - получение права на обновление информации*/
func Upd() Permission { return Update }

/*Rem - получение права на удаление информации*/
func Rem() Permission { return Remove }

/*CRRUD - Create Read ReadAll Update Remove*/
func CRRUD() []Permission {
	return []Permission{Create, Read, ReadAll, Update, Remove}
}

/*CRUD - Create Read Update Remove*/
func CRUD() []Permission {
	return []Permission{Create, Read, Update, Remove}
}

/*CRD - Create Read Remove*/
func CRD() []Permission {
	return []Permission{Create, Read, Remove}
}

/*CUD - Create Update Remove*/
func CUD() []Permission {
	return []Permission{Create, Update, Remove}
}

/*CRR - Create Read ReadAll*/
func CRR() []Permission {
	return []Permission{Create, Read, ReadAll}
}

/*UD - Update Remove*/
func UD() []Permission {
	return []Permission{Update, Remove}
}

/*RUD - Read Update Remove*/
func RUD() []Permission {
	return []Permission{Read, Update, Remove}
}

/*RR - Read ReadAll */
func RR() []Permission {
	return []Permission{Read, ReadAll}
}

/*FArticle - фича статьи*/
func FArticle() Features {
	return Article
}

/*FComment - фича комментарии*/
func FComment() Features {
	return Comment
}

/*FTag - фича тэгов*/
func FTag() Features {
	return Tag
}

/*FUser - фича пользователей*/
func FUser() Features {
	return User
}

/*FToken - фича токенов*/
func FToken() Features {
	return Token
}

const (
	/*Create - создание*/
	Create Permission = 0
	/*Read - чтение одного элемента*/
	Read Permission = 1
	/*ReadAll - чтение списка элементов*/
	ReadAll Permission = 2
	/*Update - обновление элемента по его идентификатору*/
	Update Permission = 3
	/*Remove - удаление элемента по его идентификатору*/
	Remove Permission = 4

	/*Article - фича статей*/
	Article Features = 0
	/*Comment - фича комментариев */
	Comment Features = 1
	/*Tag - фича тэгов*/
	Tag Features = 2
	/*User - фича пользователей*/
	User Features = 3
	/*Token - фича токенов*/
	Token Features = 4

	/*Auth - мидлварь аутентификации*/
	Auth MiddleWare = 0
	/*Logger - мидлварь логгирования*/
	Logger MiddleWare = 1
	/*Routes - мидлварь роутеров*/
	Routes MiddleWare = 2
)
