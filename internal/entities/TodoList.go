package entities

import "time"

type TodoList struct {
	ID         int       `json:"id"`
	UserID     int       `json:"userId"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	DeletedAt  time.Time `json:"deletedAt"`
	Percentage int       `json:"percentage"`
}
