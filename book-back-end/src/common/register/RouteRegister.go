package register

import (
	"book-back-end/src/action"
	"book-back-end/src/common/router"
	"book-back-end/src/domain/ioc"
	result "book-back-end/src/domain/models/dto"
	"book-back-end/src/domain/models/req"
	"fmt"
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
	webServer.GET(router.GET_CURRENT_USER, func(context *gin.Context) {
		action, _ := ioc.BuildUserAction()
		header := context.GetHeader("token")
		context.JSON(http.StatusOK, action.GetLoginUser(header))
	})
}

func CategoryRoute(webServer *gin.Engine) {
	webServer.POST(router.CREATE_CATEGORY, func(context *gin.Context) {
		categoryAction, _ := ioc.BuildCategoryAction()
		var req req.CategoryReq
		context.BindJSON(&req)
		userInfo, _ := context.Get(router.CONTEXT_USER)
		req.UserNo = userInfo.(*result.JwtClaims).UserNo
		context.JSON(http.StatusOK, categoryAction.CreateCategory(req))
	})
	webServer.POST(router.DEL_CATEGORY, func(context *gin.Context) {
		categoryAction, _ := ioc.BuildCategoryAction()
		var req req.CategoryReq
		context.BindJSON(&req)
		userInfo, _ := context.Get(router.CONTEXT_USER)
		req.UserNo = userInfo.(*result.JwtClaims).UserNo
		context.JSON(http.StatusOK, categoryAction.DelCategory(req))
	})
	webServer.POST(router.QUERY_CATEGORY_LIST, func(context *gin.Context) {
		categoryAction, _ := ioc.BuildCategoryAction()
		userInfo, _ := context.Get(router.CONTEXT_USER)
		userNo := userInfo.(*result.JwtClaims).UserNo
		context.JSON(http.StatusOK, categoryAction.QueryCategoryList(userNo))
	})
}

func FileRouter(webServer *gin.Engine) {
	webServer.POST(router.UPLOAD, func(context *gin.Context) {
		fileAction := action.NewFileAction()
		file := fileAction.UploadFile(context)
		context.JSON(http.StatusOK, file)
	})

	webServer.GET(router.GET_FILE, func(context *gin.Context) {
		value := context.Query("fileNo")
		fileName := action.FILE_BASE_PATH + value
		context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
		context.File(fileName)
	})
}
