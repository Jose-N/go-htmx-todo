package store

type UserStore interface {
	//GetUsers() ([]*model.User, error)
	//GetUser(id uint) (*model.User, error)
	CreateUser(firstName string, lastName string, email string, password string) (*uint, error)
	//UpdateUser(firstName string, lastName string, email string, password string) error
	//DeleteUser(id uint) error
}
