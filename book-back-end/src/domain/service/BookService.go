package service

import "book-back-end/src/domain/repo"

type BookService struct {
	repo *repo.BookRepository
}

func NewBookService(repo *repo.BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}
