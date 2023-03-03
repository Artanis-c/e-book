//go:build wireinject
// +build wireinject

package ioc

import (
	"book-back-end/src/action"
	"book-back-end/src/domain/repo"
	"book-back-end/src/domain/service"
	"github.com/google/wire"
)

func BuildUserAction() (*action.UserAction, error) {
	wire.Build(action.NewUserAction, service.NewUserService, repo.NewUserRepository, repo.NewDbConnection)
	return new(action.UserAction), nil
}

func BuildUserService() (*service.UserService, error) {
	wire.Build(service.NewUserService, repo.NewUserRepository, repo.NewDbConnection)
	return new(service.UserService), nil
}
