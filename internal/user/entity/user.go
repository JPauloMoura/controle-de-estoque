package entity

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Credentials `json:"credentials"`
	Permissions []Permission `json:"permissions"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Permission string

const (
	ADMIN  Permission = "admin"
	READER Permission = "reader"
)
