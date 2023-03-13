// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package ioc

import (
	"book-back-end/src/action"
	"book-back-end/src/domain/repo"
	"book-back-end/src/domain/service"
)

// Injectors from wire.go:

func BuildCategoryAction() (*action.CategoryAction, error) {
	db := repo.NewDbConnection()
	categoryRepository := repo.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryAction := action.NewCategoryAction(categoryService)
	return categoryAction, nil
}

func BuildCategoryService() (*service.CategoryService, error) {
	db := repo.NewDbConnection()
	categoryRepository := repo.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	return categoryService, nil
}

func BuildUserAction() (*action.UserAction, error) {
	db := repo.NewDbConnection()
	userRepository := repo.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userAction := action.NewUserAction(userService)
	return userAction, nil
}

func BuildUserService() (*service.UserService, error) {
	db := repo.NewDbConnection()
	userRepository := repo.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	return userService, nil
}

func BuildBookService() (*service.BookService, error) {
	db := repo.NewDbConnection()
	bookRepository := repo.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	return bookService, nil
}

func BuildBookAction() (*action.BookAction, error) {
	db := repo.NewDbConnection()
	bookRepository := repo.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookAction := action.NewBookAction(bookService)
	return bookAction, nil
}
