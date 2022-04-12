package response

//封装gin框架的返回

import (
	"github.com/gin-gonic/gin"
	"stream-video/code"
)

type Response gin.H

// New 新建立一个返回值
func New(code code.Code) *Response {
	return &Response{
		"code":    code,
		"message": code.GetMessage(),
	}
}

// WithData 数据返回
func (r *Response) WithData(d interface{}) *Response {
	(*r)["data"] = d
	return r
}

func (r *Response) WithToken(d interface{}) *Response {
	(*r)["token"] = d
	return r
}

// WithError 错误返回
func (r *Response) WithError(err error) *Response {
	(*r)["err_msg"] = err.Error()
	return r
}

// Return 返回 JSON 格式数据
func (r *Response) Return(c *gin.Context) {
	c.JSON(200, r)
}
