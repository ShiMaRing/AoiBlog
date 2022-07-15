package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
	"time"
)

type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}

// LimiterInter 定义限流器接口
type LimiterInter interface {
	Key(c *gin.Context) string                          //获取对应限流器的键值对名称
	GetBucket(key string) (*ratelimit.Bucket, bool)     //获取令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterInter //新增多个令牌桶
}

type Limiter struct {
	LimiterBuckets map[string]*ratelimit.Bucket
}

type MethodLimiter struct {
	*Limiter
}

// GetBucket 尝试去取桶，并尝试能否取到，每一个url作为键
func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.LimiterBuckets[key]
	return bucket, ok
}

// AddBuckets 根据指定的规则添加新的桶
func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterInter {
	for _, rule := range rules {
		if _, ok := l.LimiterBuckets[rule.Key]; !ok { //说明找不到对应的桶
			newB := ratelimit.NewBucketWithQuantum(
				rule.FillInterval, rule.Capacity, rule.Quantum)
			l.LimiterBuckets[rule.Key] = newB
		}
	}
	return l
}

func NewMethodLimiter() *MethodLimiter {
	return &MethodLimiter{
		&Limiter{
			LimiterBuckets: make(map[string]*ratelimit.Bucket),
		},
	}
}

func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return ""
	}
	return uri[:index] //返回参数之前的地址
}
