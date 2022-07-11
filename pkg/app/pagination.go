package app

import (
	"Aoi/global"
	"Aoi/pkg/convert"
	"github.com/gin-gonic/gin"
)

// GetPage 提供分页处理方法
func GetPage(c *gin.Context) int {
	query := c.Query("page")
	if query == "" {
		//如果找不到参数，默认第一页
		return 1
	}
	page := convert.StrTo(query).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	value := c.Query("size")
	if value == "" {
		return global.AppSetting.DefaultPageSize
	}
	size := convert.StrTo(value).MustInt()
	if size <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if size > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return size
}

func GetPageOffset(page, pagesize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pagesize
	}
	return result
}
