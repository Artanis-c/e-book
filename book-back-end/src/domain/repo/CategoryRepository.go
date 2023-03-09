package repo

import (
	"book-back-end/src/common"
	"book-back-end/src/domain/models"
	"book-back-end/src/domain/models/req"
	"book-back-end/src/domain/models/res"
	"gorm.io/gorm"
	"time"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (repo *CategoryRepository) CreateCategory(data req.CategoryReq) *models.Category {
	var now = time.Now()
	var saveData = &models.Category{
		UserNo:       data.UserNo,
		CategoryName: data.CategoryName,
		CategoryNo:   common.GenUniqueCode(),
		CreateTime:   &now,
	}
	tx := repo.db.Create(saveData)
	if tx.Error != nil {
		panic(tx.Error)
	}
	return saveData
}

func (repo *CategoryRepository) UpdateCategory(data req.CategoryReq) {
	tx := repo.db.Model(models.Category{}).Update(models.CategoryColumns.CategoryName, data.CategoryName)
	if tx.Error != nil {
		panic(tx.Error)
	}
}

func (repo *CategoryRepository) QueryCategoryExist(data req.CategoryReq) bool {
	var count int64
	repo.db.Model(models.Category{}).Where(" user_no=? and category_name=?", data.UserNo, data.CategoryName).Count(&count)
	return count > 0
}

func (repo *CategoryRepository) QueryCategory(categoryNo string) *models.Category {
	var res = &models.Category{}
	repo.db.Model(models.Category{}).Where(" category_no=?", categoryNo).Find(res)
	return res
}

func (repo *CategoryRepository) QueryCategoryList(userNo string) []res.CategoryRes {
	var resData []res.CategoryRes
	repo.db.Model(models.Category{}).Where("user_no=?", userNo).Find(&resData)
	return resData
}

func (repo *CategoryRepository) DelCategory(categoryNo string) {
	tx := repo.db.Delete(models.Category{}, categoryNo)
	if tx.Error != nil {
		panic(tx.Error)
	}
}
