package main

import (
	"book-back-end/src/common/register"
	"book-back-end/src/common/router"
	"github.com/gin-gonic/gin"
)

func main() {
	webServer := gin.New()
	webServer.Use(gin.Recovery())
	webServer.Use(router.AuthMiddleWare())
	webServer.Static("assets", "static/image/")
	register.UserRoute(webServer)
	register.CategoryRoute(webServer)
	register.FileRouter(webServer)
	register.BookRouter(webServer)
	err := webServer.Run(":9001")
	if err != nil {
		panic(err)
	}
}
