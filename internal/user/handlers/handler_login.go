package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/JPauloMoura/controle-de-estoque/internal/user/entity"
	"github.com/JPauloMoura/controle-de-estoque/pkg/response"
)

func (h handlerUser) Login(w http.ResponseWriter, r *http.Request) {
	var userCredentials entity.Credentials

	if err := json.NewDecoder(r.Body).Decode(&userCredentials); err != nil {
		slog.Warn("failed to decode credentials body", slog.String("error", err.Error()))
		response.Encode(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetUserByEmail(userCredentials.Email)
	if err != nil {
		slog.Warn("failed to get user", slog.String("error", err.Error()))
		response.Encode(w, err, http.StatusBadRequest)
		return
	}

	if user.Credentials.Password != userCredentials.Password {
		slog.Warn("failed to validade credentials", slog.String("error", "user or password invalid"))
		response.Encode(w, err, http.StatusForbidden)
		return
	}

	token := h.authorization.CreateToken(user.Email)
	response.Encode(w, token, http.StatusAccepted)
}
