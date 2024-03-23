package main

import (
	"github.com/Jose-N/go-htmx-todo/internal/model"
	"github.com/Jose-N/go-htmx-todo/internal/store/pgStore"
)

func main() {
	print("Hi im seed")

	fakeUser := model.User{
		FirstName: "Jose",
		LastName:  "Naylor",
		Email:     "fake@email.com",
		Password:  "pass123",
	}

	db := pgStore.ConnectDB()
	db.AutoMigrate(&model.User{})
	userStore := pgStore.UserStore{
		Db: db,
	}

	userStore.CreateUser(fakeUser.FirstName, fakeUser.LastName, fakeUser.Email, fakeUser.Password)
}
