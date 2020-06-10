package data

import (
	"errors"

	"github.com/LyonNee/grom_echo_demo/model"
	"github.com/LyonNee/grom_echo_demo/utils"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func AddUser(userinfo model.User) error {
	db, err := gorm.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/hudsondb?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	var user = model.User{}
	db.First(&user, "nickname = ?", userinfo.Nickname)

	if user != (model.User{}) {
		return errors.New("user is exist")
	}

	uid := uuid.NewV4()
	user = model.User{UUID: uid.String(), Nickname: userinfo.Nickname, Username: userinfo.Username, Password: utils.GetMD5HashCode(userinfo.Password)}
	db.Create(&user)
	return nil
}

func GetUserByUsername(username string) (model.User, error) {
	db, err := gorm.Open("mysql", "root:admin123@tcp(127.0.0.1:3306)/hudsondb?charset=utf8&parseTime=True&loc=Local")
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
