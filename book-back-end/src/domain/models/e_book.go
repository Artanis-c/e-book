package models

import "time"

// BookInfo 图书列表
type BookInfo struct {
	BookNo     string     `gorm:"primaryKey;column:book_no;type:char(16);not null" json:"bookNo"`
	BookName   string     `gorm:"column:book_name;type:varchar(100);not null" json:"bookName"` // 书名
	BarCode    string     `gorm:"column:bar_code;type:varchar(20);not null" json:"barCode"`    // 条形码
	CategoryNo string     `gorm:"column:category_no;type:char(16);not null" json:"categoryNo"` // 类型编码
	UserNo     string     `gorm:"index:i_user;column:user_no;type:char(16)" json:"userNo"`     // 用户编码
	Price      *float64   `gorm:"column:price;type:decimal(11,4)" json:"price"`                // 价格
	Image      string     `gorm:"column:image;type:varchar(200)" json:"image"`                 // 图片
	Remark     string     `gorm:"column:remark;type:varchar(200)" json:"remark"`               // 备注
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *BookInfo) TableName() string {
	return "book_info"
}

// BookInfoColumns get sql column name.获取数据库列名
var BookInfoColumns = struct {
	BookNo     string
	BookName   string
	BarCode    string
	CategoryNo string
	UserNo     string
	Price      string
	Image      string
	Remark     string
	CreateTime string
}{
	BookNo:     "book_no",
	BookName:   "book_name",
	BarCode:    "bar_code",
	CategoryNo: "category_no",
	UserNo:     "user_no",
	Price:      "price",
	Image:      "image",
	Remark:     "remark",
	CreateTime: "create_time",
}

// Category 类别表
type Category struct {
	CategoryNo   string     `gorm:"primaryKey;column:category_no;type:varchar(30);not null" json:"categoryNo"` // 类别编号
	CategoryName string     `gorm:"column:category_name;type:varchar(50);not null" json:"categoryName"`        // 类别名称
	UserNo       string     `gorm:"index:i_user;column:user_no;type:char(16)" json:"userNo"`                   // 所属人编号
	CreateTime   *time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *Category) TableName() string {
	return "category"
}

// CategoryColumns get sql column name.获取数据库列名
var CategoryColumns = struct {
	CategoryNo   string
	CategoryName string
	UserNo       string
	CreateTime   string
}{
	CategoryNo:   "category_no",
	CategoryName: "category_name",
	UserNo:       "user_no",
	CreateTime:   "create_time",
}

// Temp [...]
type Temp struct {
	IrtNo       string     `gorm:"primaryKey;column:irt_no;type:varchar(100);not null" json:"irtNo"`
	VisitRecord string     `gorm:"column:visit_record;type:varchar(1000);not null" json:"visitRecord"`
	FromKey     string     `gorm:"column:fromKey;type:varchar(50);not null" json:"fromKey"`
	CreateTime  *time.Time `gorm:"column:create_time;type:datetime" json:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *Temp) TableName() string {
	return "temp"
}

// TempColumns get sql column name.获取数据库列名
var TempColumns = struct {
	IrtNo       string
	VisitRecord string
	FromKey     string
	CreateTime  string
}{
	IrtNo:       "irt_no",
	VisitRecord: "visit_record",
	FromKey:     "fromKey",
	CreateTime:  "create_time",
}

// UserInfo 用户信息表
type UserInfo struct {
	UserNo     string     `gorm:"primaryKey;column:user_no;type:char(16);not null" json:"userNo"`
	UserName   string     `gorm:"column:user_name;type:varchar(20)" json:"userName"`
	OpenID     string     `gorm:"unique;column:open_id;type:varchar(100);not null" json:"openId"` // 微信openid
	CreateTime *time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP" json:"createTime"`
}

// TableName get sql table name.获取数据库表名
func (m *UserInfo) TableName() string {
	return "user_info"
}

// UserInfoColumns get sql column name.获取数据库列名
var UserInfoColumns = struct {
	UserNo     string
	UserName   string
	OpenID     string
	CreateTime string
}{
	UserNo:     "user_no",
	UserName:   "user_name",
	OpenID:     "open_id",
	CreateTime: "create_time",
}

// WeChatInfo [...]
type WeChatInfo struct {
	AppID     string `gorm:"primaryKey;column:app_id;type:varchar(100);not null" json:"appId"`
	AppSecret string `gorm:"column:app_secret;type:varchar(100);not null" json:"appSecret"`
}

// TableName get sql table name.获取数据库表名
func (m *WeChatInfo) TableName() string {
	return "we_chat_info"
}

// WeChatInfoColumns get sql column name.获取数据库列名
var WeChatInfoColumns = struct {
	AppID     string
	AppSecret string
}{
	AppID:     "app_id",
	AppSecret: "app_secret",
}
