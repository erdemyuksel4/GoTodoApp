
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
