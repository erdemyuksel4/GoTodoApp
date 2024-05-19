package inmemory

import (
	"todoapp/internal/entities"
	"todoapp/internal/storage"
	todolist "todoapp/internal/storage/in_memory/todolist"
	todoMsg "todoapp/internal/storage/in_memory/todomessage"
	userRepo "todoapp/internal/storage/in_memory/user"
)

type Data struct {
	TodoList    []*entities.TodoList
	TodoMessage []*entities.TodoMessage
	User        []*entities.User
}

func NewRepository(data Data) *storage.Repository {
	return &storage.Repository{
		TodoListRepository:    todolist.NewTodoListRepository(data.TodoList),
		TodoMessageRepository: todoMsg.NewTodoMessageRepository(data.TodoMessage),
		UserRepository:        userRepo.NewUserRepository(data.User),
	}
}
