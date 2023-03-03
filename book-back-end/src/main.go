package main

import (
	"book-back-end/src/common/register"
	"github.com/gin-gonic/gin"
)

func main() {
	webServer := gin.New()
	webServer.Use(gin.Recovery())
	register.UserRoute(webServer)
	err := webServer.Run(":9001")
	if err != nil {
		panic(err)
	}
}
