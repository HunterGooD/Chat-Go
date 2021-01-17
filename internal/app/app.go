package app

import (
	"database/sql"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

// AESKEY Ключ для AES
var AESKEY = make([]byte, 32)

// App структура главного приложения
type App struct {
	db    *sql.DB
	users []*User
}

// InitConfig загрузка всех необходимых настроек
func (a *App) InitConfig() {

}

// RunServer запуск сервера
func (a *App) RunServer() {

}
