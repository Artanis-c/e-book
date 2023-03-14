package repo

import (
	"book-back-end/src/domain/models"
	"book-back-end/src/domain/models/req"
	"book-back-end/src/domain/models/res"
	"gorm.io/gorm"
	"strings"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (repo *BookRepository) SaveBook(data *models.BookInfo) *models.BookInfo {
	tx := repo.db.Model(&models.BookInfo{}).Create(data)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return data
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

func (repo *BookRepository) QueryBookList(reqData *req.BookListReq) *res.PageRes {
	sql := strings.Builder{}
	coutSql := strings.Builder{}
	params := make([]interface{}, 0)
	var resData []res.BookListRes
	var count int64
	sql.WriteString("select  b.*,c.category_name from book_info as b inner join category c on b.category_no = c.category_no where 1=1")
	if reqData.BookName != "" {
		sql.WriteString(" and b.book_name =?")
		params = append(params, reqData.BookName)
	}
	if reqData.BarCode != "" {
		params = append(params, reqData.BarCode)
		sql.WriteString(" and b.bar_code =?")
	}
	coutSql.WriteString("select count(0) from (")
	coutSql.WriteString(sql.String())
	coutSql.WriteString(")")
	repo.db.Raw(sql.String(), params).Scan(&resData)
	repo.db.Raw(coutSql.String()).Scan(&count)
	return &res.PageRes{
		PageIndex: reqData.Page.PageIndex,
		PageSize:  reqData.Page.PageSize,
		Total:     count,
		Data:      resData,
	}
}
