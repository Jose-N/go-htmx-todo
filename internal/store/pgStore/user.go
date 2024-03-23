package pgStore

import (
	"github.com/Jose-N/go-htmx-todo/internal/model"
	"gorm.io/gorm"
)

type UserStore struct {
	Db *gorm.DB
}

func (s *UserStore) CreateUser(firstName string, lastName string, email string, password string) (*uint, error) {
	user := model.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}

	result := s.Db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user.ID, nil
}
