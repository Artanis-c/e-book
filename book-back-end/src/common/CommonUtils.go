package common

import (
	"book-back-end/src/domain/models"
	result "book-back-end/src/domain/models/dto"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"time"
)

var logger *zap.SugaredLogger

func init() {
	baseLogger, _ := zap.NewDevelopment()
	defer baseLogger.Sync()
	logger = baseLogger.Sugar()
	logger.Info("日志组件初始化完毕")
}

const TokenExpireDuration = time.Hour * 2

var JwtSignKey = []byte("minibook")

//GenUniqueCode 生成唯一识别码
func GenUniqueCode() string {
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 9999
	var randNum = rand.Intn(max-min+1) + min
	randStr := strconv.Itoa(randNum)
	return "10" + timeStamp + randStr
}

//GenJwtToken 生成用户令牌
func GenJwtToken(userInfo *models.UserInfo) *string {
	c := result.JwtClaims{
		OpenId: userInfo.OpenID,
		UserNo: userInfo.UserNo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "e-book",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedString, err := token.SignedString(JwtSignKey)
	if err != nil {
		logger.Error("生成JWT失败")
		panic(err)
	}
	logger.Info("生成JWT TOKEN ")
	return &signedString
}

func PassJwtToken(tokenStr string) (*result.JwtClaims, error) {
	var jwtClaims = result.JwtClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &jwtClaims, func(token *jwt.Token) (interface{}, error) {
		return JwtSignKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*result.JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
