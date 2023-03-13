package repo

import (
	"book-back-end/src/common"
	"book-back-end/src/domain/models"
	"book-back-end/src/domain/models/req"
	"book-back-end/src/domain/models/res"
	"gorm.io/gorm"
	"strings"
	"time"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (repo *BookRepository) SaveBook(data *req.SaveBookReq) *models.BookInfo {
	var now = time.Now()
	var saveData = &models.BookInfo{
		BookNo:     common.GenUniqueCode(),
		BookName:   data.BookName,
		BarCode:    data.BarCode,
		CategoryNo: data.CategoryNo,
		Price:      &data.Price,
		Image:      data.Image,
		CreateTime: &now,
	}
	tx := repo.db.Model(&models.BookInfo{}).Create(saveData)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return saveData
}

func (repo *BookRepository) UpdateBook(data *req.SaveBookReq) *models.BookInfo {
	var saveData = &models.BookInfo{
		BookName:   data.BookName,
		BarCode:    data.BarCode,
		CategoryNo: data.CategoryNo,
		Price:      &data.Price,
		Image:      data.Image,
	}
	tx := repo.db.Model(&models.BookInfo{}).Where("book_no=?", data.BookNo).Updates(saveData)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return saveData
}

func (repo *BookRepository) QueryBookList(reqData *req.BookListReq) []res.BookListRes {
	sql := strings.Builder{}
	params := make([]interface{}, 0)
	var resData []res.BookListRes
	sql.WriteString("select  b.*,c.category_name from book_info as b inner join category c on b.category_no = c.category_no where 1=1")
	if reqData.BookName != "" {
		sql.WriteString(" and b.book_name =?")
		params = append(params, reqData.BookName)
	}
	if reqData.BarCode != "" {
		params = append(params, reqData.BarCode)
		sql.WriteString(" and b.bar_code =?")
	}
	repo.db.Raw(sql.String(), params).Scan(&resData)
	return resData
}
