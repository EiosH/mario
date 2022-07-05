package middlewares

import (
	jwt "app/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		token := regexp.MustCompile(`^Bearer (.+)$`).FindStringSubmatch(authorization)[1]
		if claim, err := jwt.ParseToken(token); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "err",
				"err":     err.Error(),
			})
			c.Abort()
			return
		} else {
			fmt.Println("claim", claim.Id)
		}
	}
}
