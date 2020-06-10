package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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

func ParseJWT(tokenSrt string) (claims jwtCustomClaims, err error) {

	tokenClaims, err := jwt.ParseWithClaims(tokenSrt, &claims, func(token *jwt.Token) (interface{}, error) {
		cert := "-----BEGIN CERTIFICATE-----\n" + "secret" + "\n-----END CERTIFICATE-----"
		_, _ = jwt.ParseRSAPublicKeyFromPEM([]byte(cert))

		return claims, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(jwtCustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return claims, err
}
