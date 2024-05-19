package entities

import "time"

type TodoMessage struct {
	ID         int       `json:"id"`
	UserID     int       `json:"userId"`
	TodoListID int       `json:"todoListId"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	DeletedAt  time.Time `json:"deletedAt"`
	Content    string    `json:"content"`
	Status     bool      `json:"status"`
}
