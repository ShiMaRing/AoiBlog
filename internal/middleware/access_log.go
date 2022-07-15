package middleware

import (
	"Aoi/global"
	"Aoi/pkg/logger"
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

//实现writer接口，允许相应输出至文件
type AccessLogWriter struct {
	gin.ResponseWriter               //写出接口，需要实现
	body               *bytes.Buffer //缓冲区
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	write, err := w.body.Write(p) //向缓冲区内写入数据p，用来保存至日志中
	if err != nil {
		return write, err
	}
	return w.ResponseWriter.Write(p) //向response-writer中写入数据
}

// AccessLog 日志中间件
func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		access := &AccessLogWriter{
			ResponseWriter: ctx.Writer,
			body:           bytes.NewBufferString(""),
		}
		ctx.Writer = access

		begin := time.Now().Unix()
		ctx.Next() //继续执行下取，挂起当前handler
		end := time.Now().Unix()

		fields := logger.Fields{
			"request":  ctx.Request.PostForm.Encode(),
			"response": access.body.String(),
		}

		global.Logger.WithFields(fields).Infof(context.Background(), "access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			ctx.Request.Method, access.Status(), begin, end)

	}
}
