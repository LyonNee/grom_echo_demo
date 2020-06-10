package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	UUID     string
	Nickname string
	Username string
	Password string
}

type LoginIM struct {
	Username string
	Password string
}
