package user

import (
	"fmt"
)

// NoUserError - ошибка отсутствия пользователя
type NoUserError struct {
	id string
}

func newNoUserError(id string) NoUserError {
	return NoUserError{
		id: id,
	}
}

// Error возвращает текстовое представление ошибки
func (err NoUserError) Error() string {
	return fmt.Sprintf("no user by id '%s'", err.id)
}

// NoUserByUserNameError - ошибка отсутствия пользователя с UserName
type NoUserByUserNameError struct {
	userName string
}

func newNoUserByUserNameError(userName string) NoUserByUserNameError {
	return NoUserByUserNameError{
		userName: userName,
	}
}

// Error возвращает текстовое представление ошибки
func (err NoUserByUserNameError) Error() string {
	return fmt.Sprintf("user with username '%s' already exist", err.userName)
}

var (
	NoUserErr           = NoUserError{}
	NoUserByUserNameErr = NoUserByUserNameError{}
)
