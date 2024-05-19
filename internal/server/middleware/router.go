package middleware

import (
	"todoapp/internal/core"
	"todoapp/internal/entities"
)

type TokenHandler interface {
	Auth(httpVars core.HttpVariables, params *core.HttpAPIData)
}
type AuthorizationHandler interface {
	Verify(httpVars core.HttpVariables, params *core.HttpAPIData)
	VerifyAuthorization(user entities.User, field string, value string)
}
type Middleware struct {
	TokenHandler         TokenHandler
	AuthorizationHandler AuthorizationHandler
}
