package v1

import (
	"Aoi/global"
	"Aoi/internal/service"
	"Aoi/pkg/app"
	"Aoi/pkg/errcode"
	"context"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	//先定义req
	request := service.AuthRequest{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&request)
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetail(err.Error()))
		return
	}
	srv := service.NewService(c)
	err = srv.CheckAuth(&request)
	if err != nil {
		global.Logger.Errorf(context.Background(), "svc.CheckAuth err:%v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token, err := app.GenerateToken(request.AppKey, request.AppSecret)
	if err != nil {
		global.Logger.Errorf(context.Background(), "app.Generate err:%v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
