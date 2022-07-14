package routers

import (
	"Aoi/global"
	"Aoi/internal/service"
	"Aoi/pkg/app"
	"Aoi/pkg/convert"
	"Aoi/pkg/errcode"
	"Aoi/pkg/upload"
	"context"
	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	res := app.NewResponse(c)
	_, header, err := c.Request.FormFile("file")
	fileType := convert.StrTo(c.PostForm("type")).MustInt()
	if err != nil {
		errRsp := errcode.ErrorUploadFileFail.WithDetail(err.Error())
		res.ToResponse(errRsp)
		return
	}
	//避免不合法输入
	if header == nil || fileType <= 0 {
		res.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.NewService(c)
	uploadFile, err := svc.UploadFile(upload.FileType(fileType), nil, header)

	if err != nil {
		global.Logger.Errorf(context.Background(), "uploadFile fail %v", err)
		detail := errcode.ErrorUploadFileFail.WithDetail(err.Error())
		res.ToErrorResponse(detail)
		return
	}
	res.ToResponse(gin.H{
		"file_access_url": uploadFile.AccessUrl,
	})
}
