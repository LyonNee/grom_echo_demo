package data

import (
	"github.com/LyonNee/grom_echo_demo/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitSql() {
	db, err := gorm.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/hudsondb?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		panic("failed to connect database：" + err.Error())
	}

	isExist := db.HasTable(&model.User{})
	if !isExist {
		db.Set("gorm:table_options", "ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=1;").CreateTable(&model.User{})
	}
}
