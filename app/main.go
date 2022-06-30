package main

import (
	client "app/client"
	routers "app/routers"
)

func main() {
	client.Client()
	r := routers.SetupRouter()

	r.Run(":8888")
}
