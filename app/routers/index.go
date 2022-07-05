package routers

import (
	middlewares "app/routers/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) (*gin.Engine, error) {

	api := r.Group("/api")
	api.POST("/login", loginHandler)

	{
		api.POST("/register", registerHandler)
		api.GET("/userInfo", middlewares.Authorization(), getUserInfoHandler)
	}

	return r, nil
}
