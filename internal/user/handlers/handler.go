package handler

import (
	"net/http"

	"github.com/JPauloMoura/controle-de-estoque/internal/user/repository"
	"github.com/JPauloMoura/controle-de-estoque/pkg/auth"
)

type HandlerUser interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type handlerUser struct {
	repo          repository.UserRepository
	authorization auth.JwtAuth
}

func NewHandlerUser(repo repository.UserRepository, jwtAuth auth.JwtAuth) HandlerUser {
	return handlerUser{
		repo:          repo,
		authorization: jwtAuth,
	}
}
