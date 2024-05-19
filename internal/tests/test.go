package tests

import (
	"todoapp/internal/service"

	mockTodolist "todoapp/internal/tests/todolist"
	mockTodoMsg "todoapp/internal/tests/todomsg"
	mockUser "todoapp/internal/tests/user"
)

func NewMockService() *service.Service {
	return &service.Service{
		TodoMessageService: mockTodoMsg.NewMockTodoMessageService(),
		TodoListService:    mockTodolist.NewMockTodoListService(),
		UserService:        mockUser.NewMockUserService(),
	}
}
