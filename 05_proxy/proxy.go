package proxy

type IUser interface {
	Login(name, pwd string) error
}

type User struct {}
func (u *User) Login(name, pwd string) error {
	return nil
}

type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{user: user}
}

func (u *UserProxy) Login(name, pwd string) error {
	// todo login before

	if err := u.user.Login(name, pwd); err != nil {
		return err
	}

	// todo login after

	return nil
}