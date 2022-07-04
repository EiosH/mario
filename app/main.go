package main

import (
	client "app/client"
	routers "app/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	client.Client()

	r := gin.Default()
	routers.SetupRouter(r)

	r.Run(":8888")

}
