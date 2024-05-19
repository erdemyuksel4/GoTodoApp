package middleware

import (
	server "todoapp/internal/server"
	tokenHandler "todoapp/internal/server/middleware/token"
	"todoapp/internal/service"
)

func NewMiddleware(srv *service.Service) *server.Middleware {
	return &server.Middleware{
		TokenHandler: tokenHandler.NewTokenHandler(srv.UserService),
	}
}
