package handler

import (
	"net/http"

	"github.com/JPauloMoura/controle-de-estoque/internal/user/repository"
)

type HandlerUser interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type handlerUser struct {
	repo repository.UserRepository
}

func NewHandlerUser(repo repository.UserRepository) HandlerUser {
	return handlerUser{
		repo: repo,
	}
}
