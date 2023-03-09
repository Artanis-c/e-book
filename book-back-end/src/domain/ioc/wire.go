//go:build wireinject
// +build wireinject

package ioc

import (
	"book-back-end/src/action"
	"book-back-end/src/domain/repo"
	"book-back-end/src/domain/service"
	"github.com/google/wire"
)

func BuildCategoryAction() (*action.CategoryAction, error) {
	wire.Build(action.NewCategoryAction, service.NewCategoryService, repo.NewCategoryRepository, repo.NewDbConnection)
	return new(action.CategoryAction), nil
}

func BuildCategoryService() (*service.CategoryService, error) {
	wire.Build(service.NewCategoryService, repo.NewCategoryRepository, repo.NewDbConnection)
	return new(service.CategoryService), nil
}

func BuildUserAction() (*action.UserAction, error) {
	wire.Build(action.NewUserAction, service.NewUserService, repo.NewUserRepository, repo.NewDbConnection)
	return new(action.UserAction), nil
}

func BuildUserService() (*service.UserService, error) {
	wire.Build(service.NewUserService, repo.NewUserRepository, repo.NewDbConnection)
	return new(service.UserService), nil
}
