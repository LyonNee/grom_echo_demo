package http

import (
	"net/http"
	"strings"

	sql "github.com/LyonNee/grom_echo_demo/data"
	"github.com/LyonNee/grom_echo_demo/model"
	"github.com/LyonNee/grom_echo_demo/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	e := echo.New()

	e.Use(middleware.Logger())

	var mjwt = middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
	})

	e.POST("/register", register)
	e.POST("/login", login)
	e.POST("/my", my, mjwt)

	e.Logger.Fatal(e.Start(":8080"))
}

func register(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	err := sql.AddUser(*user)
	if err != nil {
		return c.String(http.StatusOK, "用户已存在")
	} else {
		return c.String(http.StatusOK, "组册成功")
	}
}

func login(c echo.Context) error {
	loginIM := new(model.LoginIM)
	if err := c.Bind(loginIM); err != nil {
		return err
	}

	var user = model.User{}
	user, err := sql.GetUserByUsername(loginIM.Username)

	if err != nil {
		return c.String(http.StatusOK, "用户不存在")
	} else if utils.GetMD5HashCode(loginIM.Password) != user.Password {
		return c.String(http.StatusOK, "密码错误")
	}

	// Generate encoded token and send it as response.
	t, err := utils.CreateJWT([]byte("secret"), user.UUID, user.Nickname)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func my(c echo.Context) error {
	tokenStr := c.Request().Header.Get("Authorization")
	if strings.Contains(tokenStr, "Bearer ") {
		tokenStr = tokenStr[7:]
	}

	claims, err := utils.ParseJWT(tokenStr)
	name := claims.Nickname
	//uuid:=claims.Uuid
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"name": name,
		})
	}
	return c.String(http.StatusOK, "Welcome "+name+" !")
}
