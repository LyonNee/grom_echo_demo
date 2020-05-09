package http

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	sql "grom_echo_demo/data"
	"grom_echo_demo/model"
	"grom_echo_demo/utils"
	"net/http"
)

func Start() {
	e := echo.New()

	e.Use(middleware.Logger())

	var mjwt = middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
	})

	e.POST("/register", register)
	e.POST("/login", login)
	e.GET("/my", my, mjwt)

	e.Logger.Fatal(e.Start(":8080"))
}

func register(c echo.Context) error {
	nickname := c.FormValue("nickname")
	username := c.FormValue("username")
	password := c.FormValue("password")

	err := sql.AddUser(nickname, username, password)
	if err != nil {
		return c.String(http.StatusBadRequest, "用户已存在")
	} else {
		return c.String(http.StatusOK, "组册成功")
	}
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user = model.User{}
	user, err := sql.GetUserByUsername(username)

	if err != nil {
		return c.String(http.StatusBadRequest, "用户不存在")
	} else if utils.GetMD5HashCode(password) != user.Password {
		return c.String(http.StatusBadRequest, "密码错误")
	}

	// Generate encoded token and send it as response.
	t, err := utils.CreateJWT([]byte("secret"), user.UUID, user.Nickname)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func my(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["nickname"].(string)
	return c.String(http.StatusOK, "Welcome "+name+" !")
}
