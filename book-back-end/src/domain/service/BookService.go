package service

import (
	"book-back-end/src/common"
	"book-back-end/src/domain/models"
	result "book-back-end/src/domain/models/dto"
	"book-back-end/src/domain/models/req"
	"book-back-end/src/domain/models/res"
	"book-back-end/src/domain/repo"
	"time"
)

type BookService struct {
	repo *repo.BookRepository
}

func NewBookService(repo *repo.BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (svc *BookService) SaveBook(data *req.SaveBookReq, user *result.JwtClaims) *models.BookInfo {
	var now = time.Now()
	var saveData = &models.BookInfo{
		BookNo:     common.GenUniqueCode(),
		BookName:   data.BookName,
		BarCode:    data.BarCode,
		CategoryNo: data.CategoryNo,
		Price:      &data.Price,
		Image:      data.Image,
		CreateTime: &now,
		UserNo:     user.UserNo,
	}
	return svc.repo.SaveBook(saveData)
}

func (svc *BookService) UpdateBook(data *req.SaveBookReq) *models.BookInfo {
	return svc.repo.UpdateBook(data)
}

func (svc *BookService) QueryBookList(reqData *req.BookListReq, user *result.JwtClaims) *res.PageRes {
	reqData.UserNo = user.UserNo
	return svc.repo.QueryBookList(reqData)
}
