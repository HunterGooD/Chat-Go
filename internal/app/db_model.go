package app

import (
	"gorm.io/gorm"
)

// UserDB таблица пользователей для БД
type UserDB struct {
	gorm.Model
	Login    string `gorm:"size:30;index"`
	Password string `gorm:"size:60"`
	Avatar   string `gorm:"type:text"`
	Token    string `gorm:"size:100"`
}

// MessageDB структура для  хранения сообщений
type MessageDB struct {
	gorm.Model
	Text string `gorm:"type:text"`
}

// RoomDB структура комнаты
type RoomDB struct {
	gorm.Model
	Name string `gorm:"size:20"`
	Hash string `gorm:"size:100"`
}
