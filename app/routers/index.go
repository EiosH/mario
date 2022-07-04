package routers

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

var identityKey = "id"

func SetupRouter(r *gin.Engine) (*gin.Engine, error) {

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "test zone",          //标识
		SigningAlgorithm: "HS256",              //加密算法
		Key:              []byte("secret key"), //密钥
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,   //刷新最大延长时间
		IdentityKey:      identityKey, //指定cookie的id
		PayloadFunc: func(data interface{}) jwt.MapClaims { //负载，这里可以定义返回jwt中的payload数据
			fmt.Println("PayloadFunc", data)
			//if v, ok := data.(*User); ok {
			//	return jwt.MapClaims{
			//		identityKey: v.UserName,
			//	}
			//}
			return jwt.MapClaims{}
		},
		//IdentityHandler: func(c *gin.Context) interface{} {
		//	claims := jwt.ExtractClaims(c)
		//	return &User{
		//		UserName: claims[identityKey].(string),
		//	}
		//},
		Authenticator: LoginHandler, //在这里可以写我们的登录验证逻辑
		Authorizator: func(data interface{}, c *gin.Context) bool { //当用户通过token请求受限接口时，会经过这段逻辑
			//if v, ok := data.(*User); ok && v.UserName == "admin" {
			//	return true
			//}
			//
			//return false

			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) { //错误时响应
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// 指定从哪里获取token 其格式为："<source>:<name>" 如有多个，用逗号隔开
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	api := r.Group("/api")
	api.POST("/login", authMiddleware.LoginHandler)
	api.POST("/logout", authMiddleware.LogoutHandler)
	api.POST("/refresh_token", authMiddleware.RefreshHandler)

	{
		api.POST("/register", registerHandler)
		api.GET("/userInfo", getUserInfoHandler)
	}

	r.Use(authMiddleware.MiddlewareFunc())

	return r, err
}
