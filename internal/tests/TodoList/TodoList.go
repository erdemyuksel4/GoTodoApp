package todolist_test

import (
	"log"
	"todoapp/internal/entities"
)

type MockTodoListService struct{}

func NewMockTodoListService() *MockTodoListService {
	return &MockTodoListService{}
}
func (m *MockTodoListService) GetAll() ([]*entities.TodoList, error) {

	todoLists := []*entities.TodoList{
		{ID: 1, UserID: 2, Percentage: 0},
		{ID: 2, UserID: 3, Percentage: 23},
	}
	return todoLists, nil
}

func (m *MockTodoListService) GetById(id int) (*entities.TodoList, error) {

	todoList := &entities.TodoList{ID: id, UserID: 3, Percentage: 23}
	return todoList, nil
}

func (m *MockTodoListService) Create(todoList *entities.TodoList) error {
	log.Println("Yeni oluşturuldu - id:", todoList.ID, "Percentage:", todoList.Percentage)

	return nil
}

func (m *MockTodoListService) UpdateById(id int, todoList *entities.TodoList) error {
	log.Printf("UpdateById metoduna gelen ID: %d, TodoList: %v", id, todoList)
	return nil
}

func (m *MockTodoListService) Update(todoList *entities.TodoList) error {
	log.Printf("Update metoduna gelen TodoList: %v", todoList)
	return nil
}

func (m *MockTodoListService) PatchPercentage(id int, percentage int) error {
	log.Println("durumu güncellendi - ID:", id, "Yüzde:", percentage)
	return nil
}

func (m *MockTodoListService) DeleteById(id int) error {
	log.Println("silindi - ID:", id)
	return nil
}
