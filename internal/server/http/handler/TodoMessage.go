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

type TodoMessageHandler struct {
	todoMessageService service.TodoMessageService
}

func NewTodoMessageHandler(todoMsgServ service.TodoMessageService) *TodoMessageHandler {
	return &TodoMessageHandler{todoMessageService: todoMsgServ}
}
func (t *TodoMessageHandler) GetAll(httpVars core.HttpVariables, params *core.HttpAPIData) {
	todos, err := t.todoMessageService.GetAll()
	filteredTodos := helpers.Filter(todos, func(todo *entities.TodoMessage) bool {
		return todo.DeletedAt.IsZero()
	})
	if err != nil || len(filteredTodos) <= 0 {
		httpVars.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	bytedTodos, err := json.Marshal(filteredTodos)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusOK)
	httpVars.Writer.Write(bytedTodos)
}
func (t *TodoMessageHandler) GetById(httpVars core.HttpVariables, params *core.HttpAPIData) {
	id := helpers.StrToInt(params.Query)

	todo, err := t.todoMessageService.GetById(id)
	if todo == nil || err != nil || !todo.DeletedAt.IsZero() {
		httpVars.Writer.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		return
	}


	bytedTodo, err := json.Marshal(todo)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusOK)
	httpVars.Writer.Write(bytedTodo)
}
func (t *TodoMessageHandler) Create(httpVars core.HttpVariables, params *core.HttpAPIData) {

	var todo entities.TodoMessage
	err := helpers.InterfaceToStruct(params.Body, &todo)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.todoMessageService.Create(&todo)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusCreated)
}
func (t *TodoMessageHandler) UpdateById(httpVars core.HttpVariables, params *core.HttpAPIData) {
	var todo entities.TodoMessage
	err := helpers.InterfaceToStruct(params.Body, &todo)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	id := helpers.StrToInt(params.Query)
	err = t.todoMessageService.UpdateById(id, &todo)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusOK)
}
func (t *TodoMessageHandler) Update(httpVars core.HttpVariables, params *core.HttpAPIData) {
	var todo entities.TodoMessage
	err := helpers.InterfaceToStruct(params.Body, &todo)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.todoMessageService.Update(&todo)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusOK)
}
func (t *TodoMessageHandler) DeleteById(httpVars core.HttpVariables, params *core.HttpAPIData) {
	id := helpers.StrToInt(params.Query)
	todo, err := t.todoMessageService.GetById(id)
	if err != nil {
		httpVars.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	err = t.todoMessageService.DeleteById(id)
	if err != nil {
		httpVars.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusOK)
}
func (t *TodoMessageHandler) PatchStatus(httpVars core.HttpVariables, params *core.HttpAPIData) {
	var status struct{ Status bool }
	err := helpers.InterfaceToStruct(params.Body, &status)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	id := helpers.StrToInt(params.Query)
	todo, err := t.todoMessageService.GetById(id)
	if err != nil {
		httpVars.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	err = t.todoMessageService.PatchStatus(id, status.Status)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusOK)
}
