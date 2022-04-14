package common

import (
	"context"
	"strconv"
	"stream-video/dbops"
	"time"
)

type Token struct {
	TokenKey string
}

func GetTokenKey(userId uint) *Token {
	return &Token{
		TokenKey: "UserToken:" + strconv.Itoa(int(userId)),
	}
}

// AddRedisToken 在redis当中添加对应用户的token
func (t Token) AddRedisToken(token string) error {
	expirationTime := 10800 * time.Second //过期时间 三小时
	//redis
	tokenKey := t.TokenKey //UserToken:11
	err := dbops.RDB.SetEX(context.Background(), tokenKey, token, expirationTime).Err()
	return err
}

// RedisExistToken redis是否存在对应用户的token
func (t Token) RedisExistToken() bool {
	tokenKey := t.TokenKey

	value, err := dbops.RDB.Exists(context.Background(), tokenKey).Result()
	if err != nil || value != 1 {
		return false
	}
	return true
}
