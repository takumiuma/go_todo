//go:build wireinject
// +build wireinject

package di

import (
	"practice/gateway"
	"practice/handler"
	"practice/resource"
	"practice/usecase"

	"github.com/google/wire"
)

func InitSystemHandler() *handler.SystemHandler {
	wire.Build(handler.NewSystemHandler)
	return nil
}

func InitTodoHandler() *handler.TodoHandler {
	wire.Build(
		handler.ProvideTodoHandler,
		usecase.ProvideTodoUsecase,
		gateway.ProvideTodoPort,
		resource.ProvideTodoDriver,
		resource.ProvideDatabaseConnection,
	)
	return nil
}