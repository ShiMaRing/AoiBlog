package middleware

import (
	"Aoi/pkg/app"
	"Aoi/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		//先在query路径找，找不到则在header找
		query, ok := c.GetQuery("token")
		if ok {
			token = query
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				ecode = errcode.UnauthorizedTokenError
			}
		}
		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort() //阻止剩下的中间件以及请求调用
			return
		}
		c.Next() //传递下取
	}
}
