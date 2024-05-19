package handler

import (
	"encoding/json"
	"net/http"
	core "todoapp/internal/core"
	helpers "todoapp/internal/core/helpers"
	jwtHelper "todoapp/internal/core/security/jwt"
	"todoapp/internal/entities"
	service "todoapp/internal/service"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}
func (h *UserHandler) Login(httpVars core.HttpVariables, params *core.HttpAPIData) {
	var user entities.User
	err := helpers.InterfaceToStruct(params.Body, &user)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	usr, err := h.UserService.FindByUser(user)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusUnauthorized)
		return
	}

	tokenString, err := jwtHelper.GenerateJWT(usr.UserName)
	if err != nil {
		httpVars.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	helpers.SetCookie(httpVars.Writer, httpVars.Request, usr, "user")
	bytedTokenString, err := json.Marshal(tokenString)
	if err != nil {
		http.Error(httpVars.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	httpVars.Writer.WriteHeader(http.StatusOK)
	httpVars.Writer.Write(bytedTokenString)
}
