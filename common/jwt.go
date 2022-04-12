package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"stream-video/model"
	"time"
)

//token

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// ReleaseToken 发放用户token
func ReleaseToken(user model.User) (string, error) {
	jwtKey := []byte(viper.GetString("token.jwt_key"))

	//设置token 过期时间 3小时过期
	expirationTime := time.Now().Add(3 * time.Hour)

	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),     //签发时间
			Issuer:    "stream_video",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	jwtKey := []byte(viper.GetString("token.jwt_key"))
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, e error) {
		return jwtKey, nil
	})
	return token, claims, err
}
