package router

import (
	"book-back-end/src/common"
	result "book-back-end/src/domain/models/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	WxLogin        string = "/api/user/login"
	GetCurrentUser string = "/api/user/current"
)
const (
	CreateCategory    = "/api/category/create"
	DelCategory       = "/api/category/del"
	QueryCategoryList = "/api/category/list"
)

const (
	UPLOAD  string = "/api/file/upload"
	GetFile string = "/api/file"
)

const (
	AuthHeader  string = "token"
	ContextUser string = "CONTEXT_USER"
)

const (
	AddBook       string = "/api/book/add"
	UpdateBook           = "/api/book/update"
	QueryBookList        = "api/book/list"
	QueryBook            = "/api/book/detail"
)

var WHITE_URL []string

func init() {
	WHITE_URL = make([]string, 1)
	WHITE_URL = append(WHITE_URL, WxLogin)
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		needAuth := true
		if !strings.HasPrefix(context.Request.URL.Path, "asset") {
			for _, ele := range WHITE_URL {
				if ele == context.Request.URL.Path {
					needAuth = false
					break
				}
			}
		}
		if needAuth {
			token := context.GetHeader(AuthHeader)
			if strings.HasPrefix(context.Request.URL.Path, "/assets") {
				queryToken, b := context.GetQuery("token")
				if b {
					token = queryToken
				}
			}
			if token == "" {
				context.Abort()
				context.JSON(http.StatusUnauthorized, result.Fail("无权访问"))
			} else {
				jwtToken, err := common.PassJwtToken(token)
				if err != nil {
					context.Abort()
					context.JSON(http.StatusUnauthorized, result.Fail("token 无效"))
				}
				context.Set(ContextUser, jwtToken)
				context.Next()
			}
		} else {
			context.Next()
		}
	}
}
