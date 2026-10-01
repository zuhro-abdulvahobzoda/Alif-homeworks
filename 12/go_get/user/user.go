package user

import (
	"github.com/google/uuid"
)

type User struct {
	Name string
	ID   string
}

func NewUser(name string) User {
	return User{Name: name, ID: uuid.NewString()}
}
