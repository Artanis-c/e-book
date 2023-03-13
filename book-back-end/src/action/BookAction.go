package action

import "book-back-end/src/domain/service"

type BookAction struct {
	svc *service.BookService
}

func NewBookAction(svc *service.BookService) *BookAction {
	return &BookAction{
		svc: svc,
	}
}
