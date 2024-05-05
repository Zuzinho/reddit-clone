package loginform

import "fmt"

// IncorrectLoginOrPasswordError - ошибка неправильного логина или пароля
type IncorrectLoginOrPasswordError struct {
	login string
	pass  string
}

func newIncorrectLoginOrPasswordError(login, pass string) IncorrectLoginOrPasswordError {
	return IncorrectLoginOrPasswordError{
		login: login,
		pass:  pass,
	}
}

// Error возвращает текстовое представление ошибки
func (err IncorrectLoginOrPasswordError) Error() string {
	return fmt.Sprintf("incorrect login '%s' or password '%s'", err.login, err.pass)
}

// ExistLoginFormError - ошибка занятого логина
type ExistLoginFormError struct {
	login string
}

func newExistLoginFormError(login string) ExistLoginFormError {
	return ExistLoginFormError{
		login: login,
	}
}

// Error возвращает текстовое представление ошибки
func (err ExistLoginFormError) Error() string {
	return fmt.Sprintf("user with login '%s' already exists", err.login)
}

var (
	IncorrectLoginOrPasswordErr = IncorrectLoginOrPasswordError{}
	ExistLoginFormErr           = ExistLoginFormError{}
)
