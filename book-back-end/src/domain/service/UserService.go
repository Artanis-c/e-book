package service

import (
	"book-back-end/src/common"
	"book-back-end/src/domain/models"
	"book-back-end/src/domain/models/req"
	"book-back-end/src/domain/models/res"
	"book-back-end/src/domain/repo"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type UserService struct {
	userRepo *repo.UserRepository
}

func NewUserService(repo *repo.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

//WxLogin 微信登录
func (s *UserService) WxLogin(wxLoginReq req.WxLoginReq) (*string, error) {
	session, err := s.Code2Session(wxLoginReq)
	if err != nil {
		return nil, err
	}
	var userInfo *models.UserInfo
	if s.userRepo.QueryUserInfoExist(session.OpenId) {
		userInfo = s.userRepo.QueryUserInfo(session.OpenId)
	}
	userInfo = s.userRepo.CreateUserInfo(req.CreateUserReq{OpenId: session.OpenId})
	token := common.GenJwtToken(userInfo)
	return token, nil
}

//Code2Session 置换微信OpenId
func (s *UserService) Code2Session(req req.WxLoginReq) (*res.Code2SessionRes, error) {
	wxInfo := s.userRepo.QueryWxInfo()
	url := strings.Builder{}
	url.WriteString(common.CODE_2_SESSION)
	url.WriteString("?grant_type=authorization_code")
	url.WriteString("&appid=")
	url.WriteString(wxInfo.AppID)
	url.WriteString("&secret=")
	url.WriteString(wxInfo.AppSecret)
	url.WriteString("&js_code=")
	url.WriteString(req.Code)
	reqUrl := url.String()
	client := &http.Client{}
	resp, err := client.Get(reqUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	sessionRes := res.Code2SessionRes{}
	json.Unmarshal(body, &sessionRes)
	if sessionRes.ErrMsg != "" {
		return nil, errors.New(sessionRes.ErrMsg)
	}
	return &sessionRes, nil
}

func (s *UserService) TempSave(data req.TempReq) {
	s.userRepo.TempSave(data)
}
