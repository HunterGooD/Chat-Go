package app

import "gorm.io/gorm"

type UserDB struct {
	gorm.Model
	Login    string `gorm:"size:30"`
	Password string `gorm:"size:60"`
}
