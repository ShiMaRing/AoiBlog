package middleware

import (
	"Aoi/pkg/app"
	"Aoi/pkg/errcode"
	"Aoi/pkg/limiter"
	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.LimiterInter) func(*gin.Context) {
	return func(ctx *gin.Context) {
		key := l.Key(ctx)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1) //要么拿这么多要么为0
			if count == 0 {
				response := app.NewResponse(ctx)
				response.ToErrorResponse(errcode.TooManyRequests)
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}
