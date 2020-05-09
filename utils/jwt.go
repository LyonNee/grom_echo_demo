package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Uuid     string `json:"uuid"`
	Nickname string `json:"nickname"`
}

func CreateJWT(SecretKey []byte, Uuid string, Nickname string) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 1).Unix()),
		},
		Uuid,
		Nickname,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)
	return
}

func ParseJWT(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}
