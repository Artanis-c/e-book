package req

// CreateUserReq 创建用户
type CreateUserReq struct {
	OpenId string `json:"openId"`
}

// WxLoginReq 微信登录
type WxLoginReq struct {
	ErrMsg string `json:"errMsg"`
	Code   string `json:"code"`
}

type TempReq struct {
	VisitRecordV2 string `json:"visitRecordV2"`
	IrtNo         string `json:"irtNo"`
	Formkey       string `json:"formKey"`
}
