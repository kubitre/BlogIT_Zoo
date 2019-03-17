package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	mgo "gopkg.in/mgo.v2"

	"github.com/kubitre/blog/Config"
	"github.com/kubitre/blog/Dao"
	"github.com/kubitre/blog/Routes"
)

/*Application - базовый тип, объединяющий в себе роутер, dao, db*/
type Application struct {
	Database     *mgo.Database        // прослойка бд, через которую будет работать Dao layer
	Port         int                  // порт, на котором будет запущен сервис
	LogLevel     int                  // уровень логирования (0 - на базе stdout, 1 - в файл )
	DbConfigFile string               // файлик конфигурации подключения к бд
	Routers      *Routes.RouteSetting // объединяющий роутер
}

var app = &Application{}

/*ParseCommandsFromCommandLine - парсинг аргументов командной строки*/
func (app *Application) ParseCommandsFromCommandLine() *Application {

	portNumber := flag.Int("port", 9999, "главный порт для всех хэндлеров")
	dbconfigFile := flag.String("db_config", "config.toml", "файл конфигурации подключения к бд")
	logLevel := flag.Int("log", 0, "уровень логирования, 0 - на уровне stdout, 1 - в файл")

	flag.Parse()

	app = &Application{
		Port:         *portNumber,
		DbConfigFile: *dbconfigFile,
		LogLevel:     *logLevel,
		Routers:      Routes.CreateNewRouter("/v0.1"),
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

	return app
}

func init() {
	app = app.ParseCommandsFromCommandLine().Configurating()
	// app.Routers.Setting()
}

func main() {
	log.Println("api was started on port: ", strconv.Itoa(app.Port))

	Routes.StartModeRouters(map[int][]int{
		0: []int{0, 1, 2, 3, 4},
		1: []int{0, 1, 2, 3, 4},
		2: []int{0, 1, 2, 3, 4},
		3: []int{0, 1, 2, 3, 4},
		4: []int{0, 1, 2, 3, 4},
	}, app.Routers,
		app.Database,
	)

	if err := http.ListenAndServe(":"+strconv.Itoa(int(app.Port)), app.Routers.Router); err != nil {
		log.Fatal(err)
	}
}
