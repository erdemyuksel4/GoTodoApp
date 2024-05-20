
# GoTodoApp

GoTodoApp is a simple todo list application. Users can use this application to track their todo, completed tasks.


## Installation

Clone GoTodoApp to your computer:

    git clone https://github.com/kullaniciadi/GoTodoApp.git

Navigate to the project directory:

    cd GoTodoApp

Install required dependencies:

    go mod tidy

Run the application:

    go run main.go
OR

Run .exe file

    internal.exe
## Usage

This is pre-defined in-memory database. If you wish you can use or update the internal/main.go file and adapt to your own system 

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

You have to write Authorization -jwt key given by login endpoint- to the header

Login

    Send a POST request to /api/login with the following JSON data:
```
{
    "userName": "erdem",
    "password": "$2a$12$JNs00flXnrq0FJAapqKkU.DkMIajTRwnSotRkR.skBxhV1F2lbJ2q"
}
```

Create a new task

    Send a POST request to /api/todo with the following JSON data:
```
{
    "userId": 20,
    "todoListId": 30,
    "content": "Plant a tree",
    "status": false
}
```

Get all tasks

    Send a GET request to /api/todo

Get a specific task

    Send a GET request to /api/todo/{id}

Update a task

    Send a PUT request to /api/todo with the following JSON data:
```
{
    "id": 2,
    "userId": 2,
    "todoListId": 2,
    "content": "Plant a flower",
    "status": false
}
```

Patch status

    Send a PATCH request to /api/todo/{id} with the following JSON data:
```
{
    "status": true
}
```

Delete a task

    Send a DELETE request to /api/todo/{id}

Create a new task list
    
    Send a POST request to /api/todolist with the following JSON data:
```
{
    "userId":2,
    "percentage":0
}
```
Get all task lists

    Send a GET request to /api/todolist

Get a specific task list

    Send a GET request to /api/todolist/{id}

Update a task list

    Send a PUT request to /api/todolist with the following JSON data:
```

{
    "id":2,
    "userId":2,
    "percentage":50
}
```
Patch percentage of completed

    Send a PATCH request to /api/todolist/{id} with the following JSON data:
```
{
    "percentage": 52
}
```
Delete a task list

    Send a DELETE request to /api/todolist/{id}
