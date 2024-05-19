package todolist

import (
	"errors"
	"time"

	entities "todoapp/internal/entities"
)

type TodoListRepository struct {
	TodoLists []*entities.TodoList
}

func init() {

}

func (t *TodoListRepository) GetAll() ([]*entities.TodoList, error) {
	return t.TodoLists, nil
}
func (t *TodoListRepository) GetById(id int) (*entities.TodoList, error) {

	for i, todo := range t.TodoLists {

		if todo.ID == id {

			return t.TodoLists[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func (t *TodoListRepository) UpdateById(id int, TodoList *entities.TodoList) error {
	for i, todo := range t.TodoLists {
		if todo.ID == id {
			TodoList.ModifiedAt = time.Now()
			t.TodoLists[i] = TodoList
		}
	}
	return nil
}
func (t *TodoListRepository) Update(TodoList *entities.TodoList) error {
	for i, todo := range t.TodoLists {
		if todo.ID == TodoList.ID {
			TodoList.ModifiedAt = time.Now()
			t.TodoLists[i] = TodoList
		}
	}
	return nil
}
func (t *TodoListRepository) Create(TodoList *entities.TodoList) error {
	TodoList.ID = len(t.TodoLists) + 1
	TodoList.CreatedAt = time.Now()
	t.TodoLists = append(t.TodoLists, TodoList)

	return nil
}
func (t *TodoListRepository) DeleteById(id int) error {

	for i, list := range t.TodoLists {
		if id == list.ID {
			t.TodoLists[i].DeletedAt = time.Now()
		}
	}
	return nil
}
func (t *TodoListRepository) PatchPercentage(id int, percentage int) error {
	todo, err := t.GetById(id)
	if err != nil {
		return err
	}
	todo.Percentage = percentage
	todo.ModifiedAt = time.Now()
	return nil
}

func NewTodoListRepository(todoLists []*entities.TodoList) *TodoListRepository {
	return &TodoListRepository{TodoLists: todoLists}
}
