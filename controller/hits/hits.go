package hits

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"stream-video/dbops"
	"stream-video/util"
)

const (
	UserHits = "hits" //redis 视频用户点击量 key
)

func AddVideoHitsInfo(videoId int) error {
	_, err := dbops.RDB.ZAdd(context.Background(), UserHits, &redis.Z{
		Score:  0,
		Member: strconv.Itoa(videoId),
	}).Result()
	if err != nil {
		util.Log.Error("添加数据到redis 失败 err: " + err.Error())
		return fmt.Errorf("hits_redis err: %s", err)
	}
	return nil
}
