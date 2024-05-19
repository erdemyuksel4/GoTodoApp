package main

import (
	"fmt"
	"time"

	"todoapp/internal/entities"
	http "todoapp/internal/server/http"
	"todoapp/internal/server/middleware"
	"todoapp/internal/service/domain"
	"todoapp/internal/storage/inmemory"
)

func main() {
	var db = inmemory.Data{
		User: []*entities.User{
			{ID: 1, UserName: "erdem", Password: "$2a$12$JNs00flXnrq0FJAapqKkU.DkMIajTRwnSotRkR.skBxhV1F2lbJ2q", IsAdmin: true},
			{ID: 2, UserName: "erdem2", Password: "$2a$12$rW4miHcgNElsijTzvjsxe.9B9RSJWjJTx4uN3jNybA7Epk4MiIpcy", IsAdmin: false},
			{ID: 3, UserName: "erdem3", Password: "$2a$12$7TriqNUVHqkKNuqVAvLQKOxWjvxLIN1SdSp1yWEsBqpqWeDWq/iD2", IsAdmin: false},
			{ID: 4, UserName: "erdem4", Password: "$2a$12$DALH6LBhqaNwquyyd59ZYuNP4M9rVzvq8xaNoum7ioHmela2GN3wO", IsAdmin: false},
			{ID: 5, UserName: "erdem5", Password: "$2a$12$hAYeQFOfXTuT4IkEVfQgSeZKCh19FXP.6SqFrZ3aIvx596dBD6n1.", IsAdmin: false}},
		TodoList: []*entities.TodoList{
			{ID: 1, UserID: 1, CreatedAt: time.Now(), Percentage: 0},
			{ID: 2, UserID: 2, CreatedAt: time.Now(), Percentage: 0},
			{ID: 3, UserID: 3, CreatedAt: time.Now(), Percentage: 0},
			{ID: 4, UserID: 4, CreatedAt: time.Now(), Percentage: 0},
			{ID: 5, UserID: 5, CreatedAt: time.Now(), Percentage: 0},
		},
		TodoMessage: []*entities.TodoMessage{
			{ID: 1, UserID: 1, TodoListID: 1, CreatedAt: time.Now(), Content: "Merhaba1", Status: false},
			{ID: 2, UserID: 2, TodoListID: 2, CreatedAt: time.Now(), Content: "Merhaba2", Status: false},
			{ID: 3, UserID: 3, TodoListID: 3, CreatedAt: time.Now(), Content: "Merhaba3", Status: false},
			{ID: 4, UserID: 4, TodoListID: 4, CreatedAt: time.Now(), Content: "Merhaba4", Status: false},
			{ID: 5, UserID: 5, TodoListID: 5, CreatedAt: time.Now(), Content: "Merhaba5", Status: false},
		},
	}
	repositories := inmemory.NewRepository(db)
	services := domain.NewService(repositories)
	myServer := http.NewServer(services)
	myServer.Middleware = middleware.NewMiddleware(services)
	fmt.Println("Server started on :80")
	myServer.StartServer(":80")
}
