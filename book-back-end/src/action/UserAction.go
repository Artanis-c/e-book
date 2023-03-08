package action

import (
	"book-back-end/src/common"
	"book-back-end/src/common/logutils"
	result "book-back-end/src/domain/models/dto"
	"book-back-end/src/domain/models/req"
	"book-back-end/src/domain/service"
)

type UserAction struct {
	svc *service.UserService
}

func NewUserAction(service *service.UserService) *UserAction {
	return &UserAction{
		svc: service,
	}
}

// WxLogin 用户登录
func (ac *UserAction) WxLogin(req req.WxLoginReq) result.Result {
	logutils.LogInfo("微信登录")
	logutils.LogErr(req.Code)
	loginRes, err := ac.svc.WxLogin(req)
	if err != nil {
		return result.Fail(err.Error())
	}
	return result.Success(loginRes)
}

// GetLoginUser 获取当前登录用户
func (ac *UserAction) GetLoginUser(token string) result.Result {
	jwtToken, _ := common.PassJwtToken(token)
	return result.Success(jwtToken)
}
