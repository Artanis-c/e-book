package service

import (
	"book-back-end/src/domain/models"
	"book-back-end/src/domain/models/req"
	"book-back-end/src/domain/models/res"
	"book-back-end/src/domain/repo"
	"errors"
)

type CategoryService struct {
	repo *repo.CategoryRepository
}

func NewCategoryService(repo *repo.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (svc *CategoryService) CrateCategory(reqData req.CategoryReq) (*models.Category, error) {
	exist := svc.repo.QueryCategoryExist(reqData)
	if exist {
		return nil, errors.New("已经存在相同名称的类型，不可重复添加")
	}
	category := svc.repo.CreateCategory(reqData)
	return category, nil
}

func (svc *CategoryService) DelCategory(reqData req.CategoryReq) {
	svc.repo.DelCategory(reqData.CategoryNo)
}

func (svc *CategoryService) QueryCategoryList(userNo string) []res.CategoryRes {
	list := svc.repo.QueryCategoryList(userNo)
	return list
}
