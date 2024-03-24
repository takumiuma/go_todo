package di

import (
	"practice/gateway"
	"practice/handler"
	"practice/resource"
	"practice/usecase"
)

func InitSystemHandler() *handler.SystemHandler {
	systemHandler := handler.NewSystemHandler()
	return systemHandler
}

func InitTodoHandler() *handler.TodoHandler {
	db := resource.ConnectToDatabase()
	todoDriver := resource.ProvideTodoDriver(db)
	todoPort := gateway.ProvideTodoPort(todoDriver)
	todoUsecase := usecase.ProvideTodoUsecase(todoPort)
	todoHandler := handler.ProvideTodoHandler(todoUsecase)
	return todoHandler
}