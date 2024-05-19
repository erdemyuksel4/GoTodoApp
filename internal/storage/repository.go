package storage

import (
	entities "todoapp/internal/entities"
)

type TodoListRepository interface {
	GetAll() ([]*entities.TodoList, error)
	GetById(id int) (*entities.TodoList, error)
	Update(todoMessage *entities.TodoList) error
	UpdateById(id int, todoMessage *entities.TodoList) error
	PatchPercentage(id int, percentage int) error
	Create(todoMessage *entities.TodoList) error
	DeleteById(id int) error
}
type TodoMessageRepository interface {
	GetAll() ([]*entities.TodoMessage, error)
	GetById(id int) (*entities.TodoMessage, error)
	Update(todoMessage *entities.TodoMessage) error
	UpdateById(id int, todoMessage *entities.TodoMessage) error
	PatchStatus(id int, status bool) error
	Create(todoMessage *entities.TodoMessage) error
	DeleteById(id int) error
}

type UserRepository interface {
	FindByUser(usr entities.User) (*entities.User, error)
}

type Repository struct {
	TodoListRepository    TodoListRepository
	TodoMessageRepository TodoMessageRepository
	UserRepository        UserRepository
}
