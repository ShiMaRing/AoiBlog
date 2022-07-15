package middleware

import (
	"Aoi/global"
	"Aoi/pkg/app"
	"Aoi/pkg/email"
	"Aoi/pkg/errcode"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {
	defailtMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				s := "panic recover err:%v"
				global.Logger.WithCallersFrames().Errorf(context.Background(), s, err)

				//邮件发送模块
				err2 := defailtMailer.SendEmail(global.EmailSetting.To, fmt.Sprintf("异常抛出，发生时间: %d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err))
				if err2 != nil {
					global.Logger.Panicf(context.Background(), "mail.sendmail fail with %v", err2)
				}

				app.NewResponse(ctx).ToErrorResponse(errcode.ServerError)
				ctx.Abort() //阻止继续传递
			}
		}()
		ctx.Next()
	}
}
