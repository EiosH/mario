package routers

import (
	"fmt"
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

	fmt.Println(id, password)

	if v, ok := userMap[id]; ok {
		fmt.Println(v)

		if v.Password == password {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
				"status":  0,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "密码错误",
				"status":  1,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "无该用户",
			"status":  1,
		})
	}
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
