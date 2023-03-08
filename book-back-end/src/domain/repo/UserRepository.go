package repo

import (
	"book-back-end/src/common"
	"book-back-end/src/domain/models"
	"book-back-end/src/domain/models/req"
	"gorm.io/gorm"
	"time"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) QueryUserInfoExist(openId string) bool {
	var count int64
	repo.db.Model(models.UserInfo{}).Where("open_id=?", openId).Count(&count)
	return count > 0
}

func (repo *UserRepository) CreateUserInfo(userReq req.CreateUserReq) *models.UserInfo {
	var now = time.Now()
	var user = models.UserInfo{OpenID: userReq.OpenId, CreateTime: &now, UserNo: common.GenUniqueCode()}
	repo.db.Create(&user)
	return &user
}

func (repo *UserRepository) QueryUserInfo(openId string) *models.UserInfo {
	var userInfo *models.UserInfo
	repo.db.Model(models.UserInfo{}).Where("open_id=?", openId).First(userInfo)
	return userInfo
}

func (repo *UserRepository) QueryWxInfo() *models.WeChatInfo {
	wxInfo := &models.WeChatInfo{}
	if common.APP_ID == "" {
		repo.db.Where(models.WeChatInfo{}).First(wxInfo)
		common.APP_ID = wxInfo.AppID
		common.APP_SECRET = wxInfo.AppSecret
		return wxInfo
	} else {
		wxInfo.AppID = common.APP_ID
		wxInfo.AppSecret = common.APP_SECRET
		return wxInfo
	}
}

func (repo *UserRepository) TempSave(data req.TempReq) {

	now := time.Now()
	temp := models.Temp{VisitRecord: data.VisitRecordV2, IrtNo: data.IrtNo, FromKey: data.Formkey, CreateTime: &now}
	repo.db.Create(&temp)
}
