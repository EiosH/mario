package routers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	api.POST("/login", loginHandler)
	api.POST("/logout", logoutHandler)
	api.POST("/register", registerHandler)
	api.GET("/userInfo", getUserInfoHandler)
	return r
}
