package like

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"stream-video/dbops"
	"stream-video/util"
	"time"
)

const (
	//KeyVideoLikeZetPrefix  视频点赞key
	KeyVideoLikeZetPrefix = "video:like:videoId:"
	// KeyLikeNumberZSet key是 帖子id value 是视屏的点赞数量
	KeyLikeNumberZSet = "video:like:number"
)

// IsVideoLike 判断用户是否对此视频点赞 点过 true 未点过 false
func IsVideoLike(userId int, videoId int) bool {
	//查询redis中 是否存在用户的id
	ctx := context.Background()
	key := KeyVideoLikeZetPrefix + strconv.Itoa(videoId)

	ok, _ := dbops.RDB.SIsMember(ctx, key, strconv.Itoa(userId)).Result()
	if ok {
		return true
	}
	return false
}

// DoVideoLike 为视频点赞函数 放入有序集合中
func DoVideoLike(userId int, videoId int) bool {
	ctx := context.Background()

	value := &redis.Z{
		Score:  float64(time.Now().Unix()), //点赞时间
		Member: strconv.Itoa(userId),       //用户id
	}

	key := KeyVideoLikeZetPrefix + strconv.Itoa(videoId)

	//使用redis 事务
	pipe := dbops.RDB.TxPipeline()
	//添加至redis
	pipe.ZAdd(ctx, key, value)
	//video点赞数+1
	pipe.ZIncrBy(ctx, KeyLikeNumberZSet, 1, strconv.Itoa(videoId))

	_, err := pipe.Exec(ctx)
	if err != nil {
		util.Log.Error("新增点赞失败 err: " + err.Error())
		return false
	}
	return true
}

// RemVideoLike 取消赞
func RemVideoLike(userId int, videoId int) bool {
	ctx := context.Background()
	key := KeyVideoLikeZetPrefix + strconv.Itoa(videoId)

	pipe := dbops.RDB.TxPipeline()
	pipe.ZRem(ctx, key, strconv.Itoa(userId)) //移出点赞用户池

	pipe.ZIncrBy(ctx, KeyLikeNumberZSet, -1, strconv.Itoa(videoId)) //此video点赞总数减1

	_, err := pipe.Exec(ctx)
	if err != nil {
		util.Log.Error("取消点赞失败 err: " + err.Error())
		return false
	}

	return true
}
