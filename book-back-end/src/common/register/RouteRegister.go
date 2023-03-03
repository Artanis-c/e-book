package register

import (
	"book-back-end/src/common/router"
	"book-back-end/src/domain/ioc"
	"book-back-end/src/domain/models/req"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRoute(webServer *gin.Engine) {
	webServer.POST(router.WX_LOGIN, func(context *gin.Context) {
		action, _ := ioc.BuildUserAction()
		var req req.WxLoginReq
		context.BindJSON(&req)
		context.JSON(http.StatusOK, action.WxLogin(req))
	})
}
