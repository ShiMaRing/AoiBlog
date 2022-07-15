package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeout(t time.Duration) func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		//获取请求的context上下文，并设置超时时间
		timeout, cancelFunc := context.WithTimeout(ctx.Request.Context(), t)
		defer cancelFunc()

		//传入新的context获取新的request对象并赋值
		ctx.Request = ctx.Request.WithContext(timeout)
		ctx.Next()
	}

}
