package routers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/login", loginHandler)
	r.POST("/logout", logoutHandler)
	r.POST("/register", registerHandler)
	r.GET("/userInfo", getUserInfoHandler)
	return r
}
