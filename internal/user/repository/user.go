package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/JPauloMoura/controle-de-estoque/internal/user/entity"
)

type UserRepository interface {
	CreateUser(user entity.User) (*entity.User, error)
	UpdateUser(user entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
	DeleteUser(id int) error
}

func NewUserRepository(db *sql.DB) UserRepository {
	return repository{db: db}
}

type repository struct {
	db *sql.DB
}

func (e repository) CreateUser(user entity.User) (*entity.User, error) {
	user.Id = COUNT_USER + 1
	MOCK_DATABASE_USER = append(MOCK_DATABASE_USER, user)
	return &user, nil
}
func (e repository) UpdateUser(user entity.User) error {
	return nil
}
func (e repository) GetUserByEmail(email string) (*entity.User, error) {
	fmt.Println(MOCK_DATABASE_USER)
	for _, user := range MOCK_DATABASE_USER {
		if user.Credentials.Email == email {
			return &user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (e repository) DeleteUser(id int) error {
	return nil
}

var COUNT_USER = 1
var MOCK_DATABASE_USER = []entity.User{
	{
		Id:          1,
		Name:        "jp",
		Credentials: entity.Credentials{Email: "jp@gmail.com", Password: "1234"},
	},
}
