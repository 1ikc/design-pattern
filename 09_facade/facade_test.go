package facade

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_Login(t *testing.T) {
	us := UserService{}
	user, err := us.Login(13011122233, 123456)
	assert.Nil(t, err)
	assert.Equal(t, &User{Name: "Login"}, user)
}

func TestUserService_LoginOrRegister(t *testing.T) {
	us := UserService{}
	user, err := us.LoginOrRegister(13011122233, 123456)
	assert.Nil(t, err)
	assert.Equal(t, &User{Name: "LoginOrRegister"}, user)
}

func TestUserService_Register(t *testing.T) {
	us := UserService{}
	user, err := us.Register(13011122233, 123456)
	assert.Nil(t, err)
	assert.Equal(t, &User{Name: "Register"}, user)
}