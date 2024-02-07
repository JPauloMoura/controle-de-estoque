package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/controle-de-estoque/internal/user/entity"
	"github.com/JPauloMoura/controle-de-estoque/pkg/response"
)

// antes de criar um user devemos ver se ele já existe
// também devemos criptografa a senha
func (h handlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user entity.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		slog.Warn("failed to decode user body", slog.String("error", err.Error()))
		response.Encode(w, err, http.StatusBadRequest)
		return
	}

	h.repo.CreateUser(user)
}
