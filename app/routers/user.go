package routers

import (
	jwt "app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name     string
	Id       string
	Password string
}

var MOCK_USER_TABLE = map[string]User{
	"1":    User{Name: "ezio", Id: "1", Password: "123456"},
	"test": User{Name: "test", Id: "2", Password: "test"},
}

func loginHandler(c *gin.Context) {
	userMap := MOCK_USER_TABLE
	json := User{}
	c.BindJSON(&json)

	id := json.Id
	password := json.Password
	message := "ok"
	var token string

	if v, ok := userMap[id]; ok {

		if v.Password != password {
			message = "密码错误"
		} else {
			token, _ = jwt.GenerateToken(id)
		}
	} else {
		message = "无该用户"
	}

	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data": gin.H{
			"token": token,
		},
	})

}

func registerHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "register",
	})
}

func getUserInfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
