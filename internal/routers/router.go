package routers

import (
	"Aoi/global"
	"Aoi/internal/middleware"
	v1 "Aoi/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

//提供路由

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Tracing())
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.AccessLog())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	upload := NewUpload()
	r.POST("/upload", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.GET("/auth", v1.GetAuth)
	r.Use(middleware.ContextTimeout(time.Second))

	article := v1.NewArticle()
	tag := v1.NewTag()

	r.GET("panic", func(context *gin.Context) {
		panic("hello panic")
	})

	apiv1 := r.Group("/api/v1")

	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
