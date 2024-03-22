package store

import "github.com/Jose-N/go-htmx-todo/internal/model"

type UserStore interface {
	GetUsers() ([]*model.User, error)
	GetUser(id uint) (*model.User, error)
	CreateUser(firstName string, lastName string, email string, password string) error
	UpdateUser(firstName string, lastName string, email string, password string) error
	DeleteUser(id uint) error
}
