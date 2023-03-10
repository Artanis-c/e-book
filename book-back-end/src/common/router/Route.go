package router

import (
	"book-back-end/src/common"
	result "book-back-end/src/domain/models/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	WX_LOGIN         string = "/api/user/login"
	GET_CURRENT_USER string = "/api/user/current"
)
const (
	CREATE_CATEGORY     = "/api/category/create"
	DEL_CATEGORY        = "/api/category/del"
	QUERY_CATEGORY_LIST = "/api/category/list"
)

const (
	UPLOAD   string = "/api/file/upload"
	GET_FILE string = "/api/file/read"
)

const (
	AUTH_HEADER  string = "token"
	CONTEXT_USER string = "CONTEXT_USER"
)

var WHITE_URL []string

func init() {
	WHITE_URL = make([]string, 1)
	WHITE_URL = append(WHITE_URL, WX_LOGIN)
	WHITE_URL = append(WHITE_URL, UPLOAD)
	WHITE_URL = append(WHITE_URL, GET_FILE)
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		needAuth := true
		for _, ele := range WHITE_URL {
			if ele == context.Request.URL.Path {
				needAuth = false
			}
		}
		if needAuth {
			token := context.GetHeader(AUTH_HEADER)
			if token == "" {
				context.Abort()
				context.JSON(http.StatusUnauthorized, result.Fail("无权访问"))
			} else {
				jwtToken, err := common.PassJwtToken(token)
				if err != nil {
					context.Abort()
					context.JSON(http.StatusUnauthorized, result.Fail("token 无效"))
				}
				context.Set(CONTEXT_USER, jwtToken)
				context.Next()
			}
		} else {
			context.Next()
		}
	}
}
