package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"strconv"
	"stream-video/code"
	"stream-video/controller/hits"
	"stream-video/dbops"
	"stream-video/response"
)

// Hits 记录点击量中间件
func Hits() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() //访问接口后 增加点击量

		ctx := context.Background()
		member := c.GetInt("videoId")

		_, err := dbops.RDB.ZIncrBy(ctx, hits.UserHits, 1, strconv.Itoa(member)).Result()
		if err != nil {
			response.New(code.RedisError).WithError(err).Return(c)
			return
		}
	}
}
