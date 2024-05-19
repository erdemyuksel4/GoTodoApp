package domain

import (
	"todoapp/internal/service"
	todolist "todoapp/internal/service/domain/todolist"
	todomsg "todoapp/internal/service/domain/todomessage"
	user "todoapp/internal/service/domain/user"
	"todoapp/internal/storage"
)

func NewService(repository *storage.Repository) *service.Service {
	return &service.Service{
		TodoMessageService: todomsg.NewTodoMessageService(repository.TodoMessageRepository),
		TodoListService:    todolist.NewTodoListService(repository.TodoListRepository),
		UserService:        user.NewUserService(repository.UserRepository),
	}
}
