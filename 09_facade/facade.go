package facade

type IUser interface {
	Login(phone, code int) (*User, error)
	Register(phone, code int) (*User, error)
}

type User struct {
	Name string
}

type UserService struct {}

func (u UserService) Login(phone, code int) (*User, error) {
	return &User{Name: "Login"}, nil
}

func (u UserService) Register(phone, code int) (*User, error) {
	return &User{Name: "Register"}, nil
}

func (u UserService) LoginOrRegister(phone int, code int) (*User, error) {
	return &User{Name: "LoginOrRegister"}, nil
}