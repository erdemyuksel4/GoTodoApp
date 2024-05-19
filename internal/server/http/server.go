package http

import (
	server "todoapp/internal/server"
	handlers "todoapp/internal/server/http/handler"
	"todoapp/internal/service"
)

func NewServer(service *service.Service) *server.Server {

	return &server.Server{
		TodoMessageHandler: handlers.NewTodoMessageHandler(service.TodoMessageService),
		TodoListHandler:    handlers.NewTodoListHandler(service.TodoListService),
		UserHandler:        handlers.NewUserHandler(service.UserService),
	}
}
