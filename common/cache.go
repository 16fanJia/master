package common

import (
	"context"
	"strconv"
	"stream-video/dbops"
	"stream-video/model"
	"time"
)

// UserInfo 用户信息缓存
type UserInfo struct {
	ID int
}

func NewUserInfo(userId uint) *UserInfo {
	return &UserInfo{ID: int(userId)}
}

// Key 获取Redis中 userInfo 的key值
func (u UserInfo) Key() string {
	return "UserCache:" + strconv.Itoa(u.ID)
}

// AddUserInfoToRedis 向redis添加用户信息缓存
func (u UserInfo) AddUserInfoToRedis(userInfo model.User) error {
	expirationTime := 10800 * time.Second //缓存过期时间 三小时
	//redis
	cacheKey := u.Key()

	var ctx context.Context
	//使用管道 事务 multi/exec
	pipe := dbops.RDB.TxPipeline()
	pipe.HSet(ctx, cacheKey,
		"userId", strconv.Itoa(int(userInfo.ID)),
		"name", userInfo.Name,
		"email", userInfo.Email,
		"gender", userInfo.Gender,
	)
	pipe.Expire(ctx, cacheKey, expirationTime)
	_, err := pipe.Exec(ctx)
	return err
}

// GetFromField 获取某个字段的值
func (u UserInfo) GetFromField(field string) (string, error) {
	val, err := dbops.RDB.HGet(context.Background(), u.Key(), field).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
