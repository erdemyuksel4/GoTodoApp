package server

import (
	"todoapp/internal/core"
	"todoapp/internal/entities"
)

type TodoMessageHandler interface {
	GetById(httpVars core.HttpVariables, params *core.HttpAPIData)
	GetAll(httpVars core.HttpVariables, params *core.HttpAPIData)
	Create(httpVars core.HttpVariables, params *core.HttpAPIData)
	UpdateById(httpVars core.HttpVariables, params *core.HttpAPIData)
	Update(httpVars core.HttpVariables, params *core.HttpAPIData)
	PatchStatus(httpVars core.HttpVariables, params *core.HttpAPIData)
	DeleteById(httpVars core.HttpVariables, params *core.HttpAPIData)
}
type TodoListHandler interface {
	GetById(httpVars core.HttpVariables, params *core.HttpAPIData)
	GetAll(httpVars core.HttpVariables, params *core.HttpAPIData)
	Create(httpVars core.HttpVariables, params *core.HttpAPIData)
	UpdateById(httpVars core.HttpVariables, params *core.HttpAPIData)
	Update(httpVars core.HttpVariables, params *core.HttpAPIData)
	PatchPercentage(httpVars core.HttpVariables, params *core.HttpAPIData)
	DeleteById(httpVars core.HttpVariables, params *core.HttpAPIData)
}
type UserHandler interface {
	Login(httpVars core.HttpVariables, params *core.HttpAPIData)
}
type AuthorizationHandler interface {
	Verify(httpVars core.HttpVariables, params *core.HttpAPIData)
	VerifyAuthorization(user entities.User, field string, value string)
}
type Server struct {
	TodoListHandler    TodoListHandler
	TodoMessageHandler TodoMessageHandler
	UserHandler        UserHandler
	*Middleware
}

type TokenHandler interface {
	Auth(httpVars core.HttpVariables, params *core.HttpAPIData) bool
}
type Middleware struct {
	TokenHandler TokenHandler
}
