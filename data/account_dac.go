package data

import (
	"errors"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"grom_echo_demo/model"
	"grom_echo_demo/utils"
)

func AddUser(nickname string, username string, password string) error {
	db, err := gorm.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/Hudson.DB?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	var user = model.User{}
	db.First(&user, "nickname = ?", nickname)

	if user != (model.User{}) {
		return errors.New("user is exist")
	}

	uid, err := uuid.NewV4()
	if err != nil {
		panic("create uuid failed")
	}
	user = model.User{UUID: uid.String(), Nickname: nickname, Username: username, Password: utils.GetMD5HashCode(password)}
	db.Create(&user)
	return nil
}

func GetUserByUsername(username string) (model.User, error) {
	db, err := gorm.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/Hudson.DB?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	var user = model.User{}
	db.First(&user, "Username = ?", username)

	if user == (model.User{}) {
		return model.User{}, errors.New("user is exist")
	}

	return user, nil
}
