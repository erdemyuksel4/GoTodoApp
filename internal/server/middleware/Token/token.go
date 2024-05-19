package token

import (
	"net/http"

	"todoapp/internal/core"
	jwt "todoapp/internal/core/security/jwt"
	"todoapp/internal/service"
)

type TokenHandler struct {
	UserService service.UserService
}

func NewTokenHandler(srv service.UserService) *TokenHandler {
	return &TokenHandler{UserService: srv}
}

func (t *TokenHandler) Auth(httpVars core.HttpVariables, params *core.HttpAPIData) {
	tokenString := httpVars.Request.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(httpVars.Writer, "Unauthorized", http.StatusUnauthorized)
		return
	}

	err := jwt.ValidateToken(tokenString)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusUnauthorized)
		return
	}

}
