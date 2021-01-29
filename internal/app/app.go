package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/HunterGooD/Chat-Go/pkg/crypto"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/gorilla/websocket"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

// App структура главного приложения
type App struct {
	Conf  *Config
	db    *gorm.DB
	users []*User
}

// NewApp структура приложения
func NewApp() *App {
	return &App{}
}

// InitConfig загрузка всех необходимых настроек
func (a *App) InitConfig() {
	// TODO: packr json load
	configBOX := packr.New("configs", "../../configs/")
	conf := &Config{}
	pass := &PasswordConfig{}
	data, err := configBOX.FindString("config.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(data), conf); err != nil {
		panic(err)
	}
	data, err = configBOX.FindString("passwords.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal([]byte(data), pass); err != nil {
		panic(err)
	}
	conf.Secrets = pass
	a.Conf = conf
}

// InitDB initial database
func (a *App) InitDB() {
	switch a.Conf.DB.DBType {
	case "sqlite":
		if db, err := gorm.Open(sqlite.Open(a.Conf.DB.DBName), &gorm.Config{}); err != nil {
			panic(err)
		} else {
			a.db = db
		}
	case "mysql":
		// Соединение с БД
		var dsn = ""
		if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
			panic(err)
		} else {
			a.db = db
		}
	case "postgres":
		// Соединение с БД
		var dsn = ""
		if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
			panic(err)
		} else {
			a.db = db
		}
	}
}

func (a *App) Init() {
	a.InitConfig()
	a.InitDB()

	a.db.AutoMigrate(&UserDB{})

	if os.Getenv("DEVEL") != "" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Для получения ключа AES
	aes := []byte(crypto.Base64Decode(a.Conf.Secrets.AESKey))
	fmt.Println(len(aes))
	fmt.Println(aes)
	fmt.Printf("\n%#v\n", string(aes))
	fmt.Printf("\n%#v\n", aes)
}

// RunServer запуск сервера
func (a *App) RunServer() {
	htmlFiles := packr.New("htmlFiles", "../../web/dist/")
	router := gin.Default()
	assetsBox := packr.New("assets", "../../web/dist/assets")
	router.StaticFS("/assets/", assetsBox)

	router.GET("/api", helloWorld)

	router.NoRoute(func(c *gin.Context) {
		data, err := htmlFiles.FindString("index.html")
		if err != nil {
			panic("html not found")
		}
		c.Data(http.StatusOK, "text/html;charset=utf-8", []byte(data))
	})

	router.Run(":9090")
}
