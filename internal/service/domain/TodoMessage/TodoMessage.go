package todomsg

import (
	entities "todoapp/internal/entities"
	storage "todoapp/internal/storage"
)

type TodoMessageService struct {
	todoMessageRepository storage.TodoMessageRepository
}

func NewTodoMessageService(todoMsgRepo storage.TodoMessageRepository) *TodoMessageService {
	return &TodoMessageService{todoMessageRepository: todoMsgRepo}
}

func (s *TodoMessageService) Create(todoMsg *entities.TodoMessage) error {
	err := s.todoMessageRepository.Create(todoMsg)
	if err != nil {
		return err
	}

	return nil
}
func (s *TodoMessageService) GetAll() ([]*entities.TodoMessage, error) {
	todos, err := s.todoMessageRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return todos, nil
}
func (s *TodoMessageService) GetById(id int) (*entities.TodoMessage, error) {

	todo, err := s.todoMessageRepository.GetById(id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}
func (s *TodoMessageService) UpdateById(id int, todoMsg *entities.TodoMessage) error {
	err := s.todoMessageRepository.UpdateById(id, todoMsg)
	if err != nil {
		return err
	}
	return nil
}
func (s *TodoMessageService) Update(todoMsg *entities.TodoMessage) error {
	err := s.todoMessageRepository.Update(todoMsg)
	if err != nil {
		return err
	}
	return nil
}
func (s *TodoMessageService) PatchStatus(id int, status bool) error {
	err := s.todoMessageRepository.PatchStatus(id, status)
	if err != nil {
		return err
	}
	return nil

}
func (s *TodoMessageService) DeleteById(id int) error {
	err := s.todoMessageRepository.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}
