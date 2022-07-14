package app

import (
	"Aoi/global"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Claims 定义jwt的基本属性
type Claims struct {
	AppKey    string `json:"appKey,omitempty"`
	AppSecret string `json:"appSecret,omitempty"`
	jwt.StandardClaims
}

// GetJWTSecret 获取jwt密钥
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

// GenerateToken 生成token
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expire := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    appKey,
		AppSecret: appSecret,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}
	//填充方法，类型等第一个字段信息，返回token实例
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := tokenClaims.SignedString(GetJWTSecret()) //最后拿密钥签名
	return signedString, err
}

// ParseToken 接下来实现解析代码,还能简化改写
func ParseToken(token string) (claims *Claims, err error) {
	result, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	fmt.Println(err)
	if result != nil {
		claims, ok := result.Claims.(*Claims)
		fmt.Println("ok: ", ok)
		fmt.Println("valid:", result.Valid)
		if ok && result.Valid {
			return claims, nil
		}
	}
	return nil, err
}
