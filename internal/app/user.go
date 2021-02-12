package app

import (
	"crypto/rsa"

	"github.com/gorilla/websocket"
)

// User структура пользователя подключившегося по вебсокету
type User struct {
	Name       string
	Conn       *websocket.Conn
	PrivateKey *rsa.PrivateKey
}

// NewUser создание струкутры пользователя
func (u *User) NewUser(name string, conn *websocket.Conn) *User {
	return &User{
		Name: name,
		Conn: conn,
	}
}
