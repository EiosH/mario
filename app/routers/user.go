package routers

import (
	"errors"
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

func LoginHandler(c *gin.Context) (interface{}, error) {
	userMap := MOCK_USER_TABLE
	json := User{}
	c.BindJSON(&json)

	id := json.Id
	password := json.Password
	var error error

	if v, ok := userMap[id]; ok {

		if v.Password != password {
			error = errors.New("密码错误")
		}
	} else {
		error = errors.New("无该用户")

	}

	return json, error
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
