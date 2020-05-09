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
