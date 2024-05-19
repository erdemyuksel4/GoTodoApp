package service

import (
	entities "todoapp/internal/entities"
)

type TodoListService interface {
	GetAll() ([]*entities.TodoList, error)
	GetById(id int) (*entities.TodoList, error)
	Create(todoMsg *entities.TodoList) error
	UpdateById(id int, todoMsg *entities.TodoList) error
	Update(todoMsg *entities.TodoList) error
	PatchPercentage(id int, percentage int) error
	DeleteById(id int) error
}

type TodoMessageService interface {
	GetAll() ([]*entities.TodoMessage, error)
	GetById(id int) (*entities.TodoMessage, error)
	Create(todoMsg *entities.TodoMessage) error
	UpdateById(id int, todoMsg *entities.TodoMessage) error
	Update(todoMsg *entities.TodoMessage) error
	PatchStatus(id int, status bool) error
	DeleteById(id int) error
}

/* Other interfaces */

type UserService interface {
	//Login(user *entities.User) (entities.User, error)
	FindByUser(user entities.User) (*entities.User, error)
}

// Service storage of all services.
type Service struct {
	TodoListService    TodoListService
	TodoMessageService TodoMessageService
	UserService        UserService
}
