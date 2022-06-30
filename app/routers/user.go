package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
	})
}

func logoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "logout",
	})
}

func registerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "logout",
	})
}

func getUserInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "userinfo",
	})
}
