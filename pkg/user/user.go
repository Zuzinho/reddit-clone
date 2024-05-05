package user

import "main/pkg/id"

// User - тип пользователя
type User struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
}

// NewUser возвращает экземпляр User
func NewUser(userName string) *User {
	return &User{
		ID:       id.GenerateID(),
		UserName: userName,
	}
}

// UsersRepo - интерфейс для хранения User
type UsersRepo interface {
	GetByID(userID string) (*User, error)
	GetByUserName(userName string) (*User, error)
	Create(userName string) *User
}
