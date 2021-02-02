package app

import (
	"gorm.io/gorm"
)

// UserDB таблица пользователей для БД
type UserDB struct {
	gorm.Model
	Login    string `gorm:"size:30;index"`
	Password string `gorm:"size:60"`
	Avatar   int
	Token    string      `gorm:"size:100"`
	Rooms    []RoomDB    `gorm:"many2many:user_rooms"`
	Messages []MessageDB `gorm:"foreignKey:User"`
	Posts    []PostDB    `gorm:"foreignKey:User"`
}

// MessageDB структура для  хранения сообщений
type MessageDB struct {
	gorm.Model
	Text string `gorm:"type:text"`
	Room int
	User int
}

// RoomDB структура комнаты
type RoomDB struct {
	gorm.Model
	Name       string    `gorm:"size:20"`
	Hash       string    `gorm:"size:100"`
	MessageKey MessageDB `gorm:"foreignKey:Room"`
}

// ImageDB структура для хранения картинок
type ImageDB struct {
	gorm.Model
	Source  string   `gorm:"type:text"`
	Hash    string   `gorm:"size:60;index"`
	UserKey UserDB   `gorm:"foreignKey:Avatar"`
	PostKey []PostDB `gorm:"many2many:posts_image"`
}

// PostDB для постов
type PostDB struct {
	gorm.Model
	Text string `gorm:"type:text"`
	User int
	Img  int
}
