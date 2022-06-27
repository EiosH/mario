package main

import (
	client "app/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	client.Client()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
	r.Run(":8888")
}
