package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	core "todoapp/internal/core"
	helpers "todoapp/internal/core/helpers"
	"todoapp/internal/entities"
	service "todoapp/internal/service"
)

type TodoListHandler struct {
	TodoListService service.TodoListService
}

func NewTodoListHandler(todoMsgServ service.TodoListService) *TodoListHandler {
	return &TodoListHandler{TodoListService: todoMsgServ}
}
func (t *TodoListHandler) GetAll(httpVars core.HttpVariables, params *core.HttpAPIData) {
	todoLists, err := t.TodoListService.GetAll()

	filteredTodoLists := helpers.Filter(todoLists, func(todo *entities.TodoList) bool {
		return todo.DeletedAt.IsZero()
	})
	if err != nil || len(filteredTodoLists) <= 0 {
		httpVars.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	bytedTodoLists, err := json.Marshal(filteredTodoLists)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusOK)
	httpVars.Writer.Write(bytedTodoLists)
}
func (t *TodoListHandler) GetById(httpVars core.HttpVariables, params *core.HttpAPIData) {
	id := helpers.StrToInt(params.Query)

	todoList, err := t.TodoListService.GetById(id)

	if err != nil || !todoList.DeletedAt.IsZero() {
		httpVars.Writer.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}
	bytedTodoList, err := json.Marshal(todoList)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusOK)
	httpVars.Writer.Write(bytedTodoList)
}
func (t *TodoListHandler) Create(httpVars core.HttpVariables, params *core.HttpAPIData) {
	var todoList entities.TodoList
	err := helpers.InterfaceToStruct(params.Body, &todoList)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.TodoListService.Create(&todoList)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusCreated)
}
func (t *TodoListHandler) UpdateById(httpVars core.HttpVariables, params *core.HttpAPIData) {
	var todoList entities.TodoList
	err := helpers.InterfaceToStruct(params.Body, &todoList)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	id := helpers.StrToInt(params.Query)
	err = t.TodoListService.UpdateById(id, &todoList)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusOK)
}
func (t *TodoListHandler) Update(httpVars core.HttpVariables, params *core.HttpAPIData) {
	var todoList entities.TodoList
	err := helpers.InterfaceToStruct(params.Body, &todoList)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.TodoListService.Update(&todoList)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusOK)
}
func (t *TodoListHandler) DeleteById(httpVars core.HttpVariables, params *core.HttpAPIData) {
	id := helpers.StrToInt(params.Query)
	todoList, err := t.TodoListService.GetById(id)
	if err != nil {
		httpVars.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	err = t.TodoListService.DeleteById(id)
	if err != nil {
		httpVars.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusOK)
}
func (t *TodoListHandler) PatchPercentage(httpVars core.HttpVariables, params *core.HttpAPIData) {
	var percentage struct{ percentage int }
	err := helpers.InterfaceToStruct(params.Body, &percentage)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	id := helpers.StrToInt(params.Query)
	todoList, err := t.TodoListService.GetById(id)
	if err != nil {
		httpVars.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	err = t.TodoListService.PatchPercentage(id, percentage.percentage)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusOK)
}
