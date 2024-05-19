package server

import (
	"net/http"
	core "todoapp/internal/core"
	//mw "todoapp/internal/server/middleware"
)

var UrlPattern core.UrlPattern

func (s *Server) StartServer(addr string) {
	mux := http.NewServeMux()

	UrlPattern.Order = []string{"Controller", "Action", "Query"}

	core.Route("POST /api/login", s.UserHandler.Login, UrlPattern, mux)
	core.Route("GET /api/todo", s.TodoMessageHandler.GetAll, UrlPattern, mux, s.TokenHandler.Auth)
	core.Route("GET /api/todo/", s.TodoMessageHandler.GetById, UrlPattern, mux, s.TokenHandler.Auth)
	core.Route("POST /api/todo", s.TodoMessageHandler.Create, UrlPattern, mux, s.TokenHandler.Auth)
	core.Route("PUT /api/todo/", s.TodoMessageHandler.UpdateById, UrlPattern, mux, s.TokenHandler.Auth)
	core.Route("PUT /api/todo", s.TodoMessageHandler.Update, UrlPattern, mux, s.TokenHandler.Auth)
	core.Route("PATCH /api/todo/", s.TodoMessageHandler.PatchStatus, UrlPattern, mux, s.TokenHandler.Auth)
	core.Route("DELETE /api/todo/", s.TodoMessageHandler.DeleteById, UrlPattern, mux, s.TokenHandler.Auth)

	core.Route("GET /api/todolist", s.TodoListHandler.GetAll, UrlPattern, mux, s.TokenHandler.Auth)
	core.Route("GET /api/todolist/", s.TodoListHandler.GetById, UrlPattern, mux, s.TokenHandler.Auth)
	core.Route("POST /api/todolist", s.TodoListHandler.Create, UrlPattern, mux, s.TokenHandler.Auth)
	core.Route("PUT /api/todolist/", s.TodoListHandler.UpdateById, UrlPattern, mux, s.TokenHandler.Auth)
	core.Route("PUT /api/todolist", s.TodoListHandler.Update, UrlPattern, mux, s.TokenHandler.Auth)
	core.Route("PATCH /api/todolist/", s.TodoListHandler.PatchPercentage, UrlPattern, mux, s.TokenHandler.Auth)
	core.Route("DELETE /api/todolist/", s.TodoListHandler.DeleteById, UrlPattern, mux, s.TokenHandler.Auth)
	http.ListenAndServe(addr, mux)
}
