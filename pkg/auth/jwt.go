package auth

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/JPauloMoura/controle-de-estoque/pkg/response"
	"github.com/golang-jwt/jwt/v5"
)

type JwtAuth interface {
	CreateToken(email string) string
	MiddlewareAuth(next http.Handler) http.Handler
}
type jwtAuth struct {
	JwtKey string
}

func NewJwtAuth(key string) JwtAuth {
	return jwtAuth{
		JwtKey: key,
	}
}

func (j jwtAuth) CreateToken(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss": "login",
			"sub": email,
			"exp": time.Now().Add(time.Minute * 5).Unix(),
		})

	strToken, err := token.SignedString([]byte(j.JwtKey))
	if err != nil {
		panic(err)
	}

	return "Bearer " + strToken
}

func (j jwtAuth) MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lengthBearer := len("Bearer ")

		bearerToken := r.Header.Get("Authorization")
		if bearerToken == "" || len(bearerToken) < lengthBearer {
			slog.Warn("failed to get token")
			response.Encode(w, errors.New("authorization header is required"), http.StatusBadRequest)
			return
		}

		token, err := parseToken(bearerToken[lengthBearer:], j.JwtKey)
		if err != nil || !token.Valid {
			slog.Warn("failed to validate token", slog.String("error", err.Error()))
			response.Encode(w, errors.New("invalid token"), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func parseToken(token string, key string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("signing method invalid")
		}

		return []byte(key), nil
	})
}
