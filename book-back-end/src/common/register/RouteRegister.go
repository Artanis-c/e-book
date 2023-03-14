package register

import (
	"book-back-end/src/action"
	"book-back-end/src/common/router"
	"book-back-end/src/domain/ioc"
	result "book-back-end/src/domain/models/dto"
	"book-back-end/src/domain/models/req"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRoute(webServer *gin.Engine) {
	webServer.POST(router.WxLogin, func(context *gin.Context) {
		action, _ := ioc.BuildUserAction()
		var req req.WxLoginReq
		context.BindJSON(&req)
		context.JSON(http.StatusOK, action.WxLogin(req))
	})
	webServer.GET(router.GetCurrentUser, func(context *gin.Context) {
		action, _ := ioc.BuildUserAction()
		header := context.GetHeader("token")
		context.JSON(http.StatusOK, action.GetLoginUser(header))
	})
}

func CategoryRoute(webServer *gin.Engine) {
	webServer.POST(router.CreateCategory, func(context *gin.Context) {
		categoryAction, _ := ioc.BuildCategoryAction()
		var req req.CategoryReq
		context.BindJSON(&req)
		userInfo, _ := context.Get(router.ContextUser)
		req.UserNo = userInfo.(*result.JwtClaims).UserNo
		context.JSON(http.StatusOK, categoryAction.CreateCategory(req))
	})
	webServer.POST(router.DelCategory, func(context *gin.Context) {
		categoryAction, _ := ioc.BuildCategoryAction()
		var req req.CategoryReq
		context.BindJSON(&req)
		userInfo, _ := context.Get(router.ContextUser)
		req.UserNo = userInfo.(*result.JwtClaims).UserNo
		context.JSON(http.StatusOK, categoryAction.DelCategory(req))
	})
	webServer.POST(router.QueryCategoryList, func(context *gin.Context) {
		categoryAction, _ := ioc.BuildCategoryAction()
		userInfo, _ := context.Get(router.ContextUser)
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
}

func BookRouter(webServer *gin.Engine) {
	webServer.POST(router.AddBook, func(context *gin.Context) {
		bookAction, _ := ioc.BuildBookAction()
		req := req.SaveBookReq{}
		context.BindJSON(&req)
		userInfo, _ := context.Get(router.ContextUser)
		context.JSON(http.StatusOK, bookAction.SaveBook(&req, userInfo.(*result.JwtClaims)))
	})

	webServer.POST(router.UpdateBook, func(context *gin.Context) {
		bookAction, _ := ioc.BuildBookAction()
		req := req.SaveBookReq{}
		context.BindJSON(&req)
		context.JSON(http.StatusOK, bookAction.UpdateBook(&req))
	})

	webServer.POST(router.QueryBookList, func(context *gin.Context) {
		bookAction, _ := ioc.BuildBookAction()
		req := req.BookListReq{}
		context.BindJSON(&req)
		userInfo, _ := context.Get(router.ContextUser)
		context.JSON(http.StatusOK, bookAction.QueryBookList(&req, userInfo.(*result.JwtClaims)))
	})
}
