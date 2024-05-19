package todolist

import (
	entities "todoapp/internal/entities"
	storage "todoapp/internal/storage"
)

type TodoListService struct {
	TodoListRepository storage.TodoListRepository
}

func NewTodoListService(todoMsgRepo storage.TodoListRepository) *TodoListService {
	return &TodoListService{TodoListRepository: todoMsgRepo}
}

func (s *TodoListService) Create(todoMsg *entities.TodoList) error {
	err := s.TodoListRepository.Create(todoMsg)
	if err != nil {
		return err
	}

	return nil
}
func (s *TodoListService) GetAll() ([]*entities.TodoList, error) {
	todos, err := s.TodoListRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return todos, nil
}
func (s *TodoListService) GetById(id int) (*entities.TodoList, error) {

	todo, err := s.TodoListRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (s *TodoListService) UpdateById(id int, todoMsg *entities.TodoList) error {
	err := s.TodoListRepository.UpdateById(id, todoMsg)
	if err != nil {
		return err
	}
	return nil
}
func (s *TodoListService) Update(todoMsg *entities.TodoList) error {
	err := s.TodoListRepository.Update(todoMsg)
	if err != nil {
		return err
	}
	return nil
}
func (s *TodoListService) PatchPercentage(id int, percentage int) error {
	err := s.TodoListRepository.PatchPercentage(id, percentage)
	if err != nil {
		return err
	}
	return nil

}
func (s *TodoListService) DeleteById(id int) error {
	err := s.TodoListRepository.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}
