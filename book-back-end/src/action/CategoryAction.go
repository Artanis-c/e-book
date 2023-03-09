package action

import (
	result "book-back-end/src/domain/models/dto"
	"book-back-end/src/domain/models/req"
	"book-back-end/src/domain/service"
)

type CategoryAction struct {
	svc *service.CategoryService
}

func NewCategoryAction(svc *service.CategoryService) *CategoryAction {
	return &CategoryAction{
		svc: svc,
	}
}

func (c *CategoryAction) CreateCategory(reqData req.CategoryReq) result.Result {
	category, err := c.svc.CrateCategory(reqData)
	if err != nil {
		return result.Fail(err.Error())
	}
	return result.Success(category)
}

func (c *CategoryAction) DelCategory(reqData req.CategoryReq) result.Result {
	c.svc.DelCategory(reqData)
	return result.Success(reqData)
}

func (c *CategoryAction) QueryCategoryList(userNo string) result.Result {
	list := c.svc.QueryCategoryList(userNo)
	return result.Success(list)
}
