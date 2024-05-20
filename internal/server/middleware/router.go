package middleware

import (
	"todoapp/internal/core"
)

type TokenHandler interface {
	Auth(httpVars core.HttpVariables, params *core.HttpAPIData) bool
}
type Middleware struct {
	TokenHandler TokenHandler
}
