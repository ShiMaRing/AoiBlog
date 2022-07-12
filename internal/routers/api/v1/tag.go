package v1

import (
	"Aoi/global"
	"Aoi/internal/service"
	"Aoi/pkg/app"
	"Aoi/pkg/convert"
	"Aoi/pkg/errcode"
	"context"
	"github.com/gin-gonic/gin"
)

//定义Tag应当暴露的接口

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	req := service.TagReq{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&req)
	if err != nil {
		global.Logger.Errorf(context.Background(), "bind fail with %v", err)
		detail := errcode.InvalidParams.WithDetail(err.Error())
		response.ToErrorResponse(detail)
		return
	}
	newService := service.NewService(c)
	list, err := newService.GetList(&req)
	if err != nil {
		global.Logger.Errorf(context.Background(), "get tag list fail with %v", err)
		detail := errcode.InvalidParams.WithDetail(err.Error())
		response.ToErrorResponse(detail)
		return
	}
	response.ToResponseList(list)
}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	req := service.TagReq{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&req)
	if err != nil {
		global.Logger.Errorf(context.Background(), "bind fail with %v", err)
		detail := errcode.InvalidParams.WithDetail(err.Error())
		response.ToErrorResponse(detail)
		return
	}
	newService := service.NewService(c)
	err = newService.CreateTag(&req)
	if err != nil {
		global.Logger.Errorf(context.Background(), "create tag list fail with %v", err)
		detail := errcode.InvalidParams.WithDetail(err.Error())
		response.ToErrorResponse(detail)
		return
	}
	response.ToResponse("create success!")
}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Tag "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	req := service.TagReq{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&req)
	req.ID = convert.StrTo(c.Param("id")).MustUInt32()
	if err != nil {
		global.Logger.Errorf(context.Background(), "bind fail with %v", err)
		detail := errcode.InvalidParams.WithDetail(err.Error())
		response.ToErrorResponse(detail)
		return
	}
	newService := service.NewService(c)
	err = newService.UpdateTag(&req)
	if err != nil {
		global.Logger.Errorf(context.Background(), "Update tag list fail with %v", err)
		detail := errcode.InvalidParams.WithDetail(err.Error())
		response.ToErrorResponse(detail)
		return
	}
	response.ToResponse("Update success!")

}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	req := service.TagReq{}
	response := app.NewResponse(c)
	err := c.ShouldBind(&req)
	req.ID = convert.StrTo(c.Param("id")).MustUInt32()
	if err != nil {
		global.Logger.Errorf(context.Background(), "bind fail with %v", err)
		detail := errcode.InvalidParams.WithDetail(err.Error())
		response.ToErrorResponse(detail)
		return
	}
	newService := service.NewService(c)
	err = newService.DeleteTag(&req)
	if err != nil {
		global.Logger.Errorf(context.Background(), "Delete tag list fail with %v", err)
		detail := errcode.InvalidParams.WithDetail(err.Error())
		response.ToErrorResponse(detail)
		return
	}
	response.ToResponse("Delete success!")
}
