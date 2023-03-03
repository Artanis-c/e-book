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