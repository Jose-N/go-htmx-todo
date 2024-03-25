package pgStore

import (
	"fmt"

	"github.com/Jose-N/go-htmx-todo/internal/model"
	"golang.org/x/crypto/bcrypt"
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
		fmt.Println(result.Error)
		return nil, result.Error
	}

	return &user.ID, nil
}

func (s *UserStore) FindByEmailAndPassword(email string, password string) (*model.User, error) {
	user := model.User{}

	err := s.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &user, nil
}
