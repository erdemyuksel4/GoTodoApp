package todomessage

import (
	"errors"
	"time"

	entities "todoapp/internal/entities"
)

type TodoMessageRepository struct {
	TodoMessages []*entities.TodoMessage
}

func init() {

}
func NewTodoMessageRepository(todos []*entities.TodoMessage) *TodoMessageRepository {
	return &TodoMessageRepository{TodoMessages: todos}
}

func (t *TodoMessageRepository) GetAll() ([]*entities.TodoMessage, error) {
	return t.TodoMessages, nil
}
func (t *TodoMessageRepository) GetById(id int) (*entities.TodoMessage, error) {

	for i, todo := range t.TodoMessages {

		if todo.ID == id {

			return t.TodoMessages[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func (t *TodoMessageRepository) UpdateById(id int, todoMessage *entities.TodoMessage) error {
	for i, todo := range t.TodoMessages {
		if todo.ID == id {
			todoMessage.ModifiedAt = time.Now()
			t.TodoMessages[i] = todoMessage
		}
	}
	return nil
}
func (t *TodoMessageRepository) Update(todoMessage *entities.TodoMessage) error {
	for i, todo := range t.TodoMessages {
		if todo.ID == todoMessage.ID {
			todoMessage.ModifiedAt = time.Now()
			t.TodoMessages[i] = todoMessage
		}
	}
	return nil
}
func (t *TodoMessageRepository) Create(todoMessage *entities.TodoMessage) error {
	todoMessage.ID = len(t.TodoMessages) + 1
	todoMessage.CreatedAt = time.Now()
	t.TodoMessages = append(t.TodoMessages, todoMessage)

	return nil
}
func (t *TodoMessageRepository) DeleteById(id int) error {

	for i, todo := range t.TodoMessages {
		if id == todo.ID {
			t.TodoMessages[i].DeletedAt = time.Now()
		}
	}
	return nil
}
func (t *TodoMessageRepository) PatchStatus(id int, status bool) error {
	todo, err := t.GetById(id)
	if err != nil {
		return err
	}
	todo.Status = status
	todo.ModifiedAt = time.Now()
	return nil
}
