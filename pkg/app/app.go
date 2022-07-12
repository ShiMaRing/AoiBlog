package app

import (
	"Aoi/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
}

type Pager struct {
	Page int `json:"page,omitempty"`
	Size int `json:"size,omitempty"`
	Data any `json:"data"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}
func (r *Response) ToResponseList(list interface{}) {
	r.Ctx.JSON(http.StatusOK, Pager{
		Page: GetPage(r.Ctx),
		Size: GetPageSize(r.Ctx),
		Data: list,
	})
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code": err.Code(), "msg": err.Msg()}
	if len(err.Details()) > 0 {
		response["details"] = err.Details()
	}
	r.Ctx.JSON(err.StatusCode(), response)

}
