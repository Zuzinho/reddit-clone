package user

import "main/pkg/id"

type User struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
}

func NewUser(userName string) *User {
	return &User{
		ID:       id.GenerateID(),
		UserName: userName,
	}
}

type UsersRepo interface {
	GetByID(userID string) (*User, error)
	GetByUserName(userName string) (*User, error)
	Create(userName string) *User
}
