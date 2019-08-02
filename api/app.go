package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/rs/cors"

	"gopkg.in/mgo.v2"

	"blog_module/Config"
	"blog_module/Dao"
	. "blog_module/Routes"
)

type (
	/*Application - базовый тип, объединяющий в себе роутер, dao, db*/
	Application struct {
		Database     *mgo.Database // прослойка бд, через которую будет работать Dao layer
		Port         int           // порт, на котором будет запущен сервис
		LogLevel     int           // уровень логирования (0 - на базе stdout, 1 - в файл )
		DbConfigFile string        // файлик конфигурации подключения к бд
		Routers      *RouteSetting // объединяющий роутер
	}
)

var app = &Application{}

/*ParseCommandsFromCommandLine - парсинг аргументов командной строки*/
func (app *Application) ParseCommandsFromCommandLine() *Application {

	portNumber := flag.Int("port", 9997, "главный порт для всех хэндлеров")
	dbconfigFile := flag.String("db_config", "config.toml", "файл конфигурации подключения к бд")
	logLevel := flag.Int("log", 0, "уровень логирования, 0 - на уровне stdout, 1 - в файл")

	flag.Parse()

	app = &Application{
		Port:         *portNumber,
		DbConfigFile: *dbconfigFile,
		LogLevel:     *logLevel,
	}

	return app
}

/*Configurating - конифгурирование бд, роутеров*/
func (app *Application) Configurating() *Application {
	bdconfig := Config.Configuration{}
	if err := bdconfig.Read(); err != nil {
		return nil
	}

	bdRouter := Dao.SettingDao{
		Server:   bdconfig.Server,
		Database: bdconfig.Database,
	}
	bdRouter.Connect()
	app.Database = bdRouter.Db
	app.Routers = CreateNewRouter("/v0.1", app.Database)

	return app
}

func init() {
	app = app.ParseCommandsFromCommandLine().Configurating()
}

func main() {
	log.Println("api was started on port: ", strconv.Itoa(app.Port))

	StartModeRouters(map[Features]map[MiddleWare][]Permission{
		FArticle(): {
			MRouter(): RR(),
			MAuth():   CUD(),
		},
		FComment(): {
			MRouter(): RR(),
			MAuth():   CUD(),
		},
		FTag(): {
			MRouter(): RR(),
			MAuth():   CUD(),
		},
		FUser(): {
			MRouter(): CRR(),
			MAuth():   UD(),
		},
		FToken(): {
			MRouter(): []Permission{Cre()},
			MAuth():   RUD(),
		},
	}, app.Routers,
		app.Database,
	)

	handler := cors.New(
		cors.Options{
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders: []string{"Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},
			AllowedOrigins: []string{"http://localhost:3000", "http://192.168.100.3:3000", "http://localhost:3001"},
			Debug:          true,
		}).Handler(app.Routers.Router)

	if err := http.ListenAndServe(":"+strconv.Itoa(int(app.Port)), handler); err != nil {
		log.Fatal(err)
	}
}
