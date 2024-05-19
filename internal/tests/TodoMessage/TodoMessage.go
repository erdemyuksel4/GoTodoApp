package todomsg_test

import (
	"log"
	"todoapp/internal/entities"
)

type MockTodoMessageService struct{}

func NewMockTodoMessageService() *MockTodoMessageService {
	return &MockTodoMessageService{}
}
func (m *MockTodoMessageService) PatchStatus(id int, status bool) error {

	log.Println("durumu güncellendi - ID:", id, "Durum:", status)
	return nil
}
func (m *MockTodoMessageService) Create(todo *entities.TodoMessage) error {

	log.Println("Yeni oluşturuldu - id:", todo.ID, "İçerik:", todo.Content, "status:", todo.Status)
	return nil
}
func (m *MockTodoMessageService) DeleteById(id int) error {

	log.Println("silindi - ID:", id)
	return nil
}
func (m *MockTodoMessageService) GetAll() ([]*entities.TodoMessage, error) {
	return []*entities.TodoMessage{
		{ID: 1, Content: "1", Status: true},
		{ID: 2, Content: "2", Status: false},
		{ID: 3, Content: "3", Status: true},
	}, nil
}
func (m *MockTodoMessageService) GetById(id int) (*entities.TodoMessage, error) {
	return &entities.TodoMessage{ID: 1, Content: "1", Status: true}, nil
}
func (m *MockTodoMessageService) Update(message *entities.TodoMessage) error {
	log.Printf("Update metoduna gelen TodoList: %v", message)
	return nil
}
func (m *MockTodoMessageService) UpdateById(id int, message *entities.TodoMessage) error {
	log.Printf("UpdateById metoduna gelen ID: %d, TodoList: %v", id, message)
	return nil
}
