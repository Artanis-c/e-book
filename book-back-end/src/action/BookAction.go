package action

import (
	result "book-back-end/src/domain/models/dto"
	"book-back-end/src/domain/models/req"
	"book-back-end/src/domain/service"
)

type BookAction struct {
	svc *service.BookService
}

func NewBookAction(svc *service.BookService) *BookAction {
	return &BookAction{
		svc: svc,
	}
}

func (ac *BookAction) SaveBook(data *req.SaveBookReq, user *result.JwtClaims) result.Result {
	book := ac.svc.SaveBook(data, user)
	return result.Success(book)
}

func (ac *BookAction) UpdateBook(data *req.SaveBookReq) result.Result {
	book := ac.svc.UpdateBook(data)
	return result.Success(book)
}

func (ac *BookAction) QueryBookList(reqData *req.BookListReq, user *result.JwtClaims) result.Result {
	list := ac.svc.QueryBookList(reqData, user)
	return result.Success(list)
}
