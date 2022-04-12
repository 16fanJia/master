package common

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"stream-video/dbops"
	"stream-video/model"
	"time"
)

// AddRedisToken 在redis当中添加对应用户的token
func AddRedisToken(token string, userId uint) error {
	expirationTime := 10800 * time.Second //过期时间 三小时
	//redis
	tokenKey := "UserToken:" + strconv.Itoa(int(userId)) //UserToken:11
	err := dbops.RDB.SetEX(context.Background(), tokenKey, token, expirationTime).Err()
	return err
}

// GetRedisToken redis是否存在对应用户的token
func GetRedisToken(userId uint) bool {
	tokenKey := "UserToken:" + strconv.Itoa(int(userId)) //UserToken:11

	_, err := dbops.RDB.Get(context.Background(), tokenKey).Result()
	if err == redis.Nil {
		return false
	}
	return true
}

// UserInfo 用户信息缓存
type UserInfo struct {
	ID int
}

func NewUserInfo(userId uint) *UserInfo {
	return &UserInfo{ID: int(userId)}
}

// AddUserInfoToRedis 向redis添加用户信息缓存
func (u UserInfo) AddUserInfoToRedis(userInfo model.User) error {
	expirationTime := 10800 * time.Second //缓存过期时间 三小时
	//redis
	cacheKey := "UserCache:" + strconv.Itoa(int(userInfo.ID)) //UserToken:11

	errHst := dbops.RDB.HSet(context.Background(), cacheKey,
		"userId", strconv.Itoa(int(userInfo.ID)),
		"name", userInfo.Name,
		"email", userInfo.Email,
		"gender", userInfo.Gender,
	).Err()
	if errHst != nil {
		return errHst
	}

	err := dbops.RDB.Expire(context.Background(), cacheKey, expirationTime).Err()
	return err
}
